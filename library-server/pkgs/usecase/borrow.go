package usecase

import (
	"context"
	"fmt"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"
	"libstack/pkgs/util"

	"go.uber.org/zap"
)

// TODO(mchenryc): quit eating cause of internal errors

// Borrow will create a loan of the specified title for the provided user.
//
// Use Case:
// 	* (1) Only patrons can borrow titles
// 	* (2) Only valid titles can be borrowed
// 	* (3) Only titles in stock can be borrowed
//
// Implemtation Details:
// 	* (1) -> `"unauthorized"`
// 	* (2) -> `"isbn_not_found"`
// 	* (3) -> `"hold_instead"`
// 	* (unexpected behavior) -> `"internal_error"`
func (i *Interactor) Borrow(ctx context.Context, isbn string, user model.User) (model.Loan, *UseCaseError) {
	var emptyLoan model.Loan
	trace, _ := util.Uuid()
	logger := logging.From(ctx).With(
		zap.String("method", "borrow"),
		zap.String("user", user.Username),
		zap.String("isbn", isbn),
		zap.String("trace-id", trace),
	)
	logger.Info("start")
	if !user.IsPatron() {
		logger.Info("unauthorized")
		return emptyLoan, unauthorizedError("user must be a patron to borrow a title")
	}

	titleCh := make(chan Result[model.Title])
	countCh := make(chan Result[int])

	go func() {
		title, err := i.TitleRW.GetByIsbn(ctx, isbn)
		titleCh <- Result[model.Title]{Data: title, Err: err}
		close(titleCh)
	}()
	go func() {
		count, err := i.LoanRW.Count(ctx, isbn)
		countCh <- Result[int]{Data: count, Err: err}
		close(countCh)
	}()

	titleResult := <-titleCh
	title, err := titleResult.Unwrap()
	if err != nil {
		logger.Info("isbn not found")
		return emptyLoan, notFoundError(fmt.Sprintf("no title found with isbn %q", isbn))
	}

	countResult := <-countCh
	count, err := countResult.Unwrap()
	if err != nil {
		logger.Info("unable to load number of loans")
		msg := fmt.Sprintf("unable to load number of loans for title with isbn %q", isbn)
		return emptyLoan, internalError(msg, err)
	}

	// >= instead of == incase number of loans gets into a bad state
	if count >= title.Stock {
		logger.Info("title not available")
		return emptyLoan, holdInsteadError("title not available")
	}

	loan, err := i.LoanRW.Add(ctx, isbn, user)
	if err != nil {
		logger.Info("failed to add loan")
		return emptyLoan, internalError("failed to add loan", err)
	}

	logger.Info("success")
	return loan, nil
}

type Result[T any] struct {
	Data T
	Err  error
}

func (r *Result[T]) Unwrap() (T, error) {
	return r.Data, r.Err
}
