package usecase

import (
	"context"
	"fmt"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"

	"go.uber.org/zap"
)

type BorrowError struct {
	Kind    string
	message string
}

func (e *BorrowError) Error() string {
	return fmt.Sprintf("%s: %s", e.Kind, e.message)
}

func internalError(message string) *BorrowError {
	return &BorrowError{Kind: "internal_error", message: message}
}
func unauthorizedError(message string) *BorrowError {
	return &BorrowError{Kind: "unauthorized", message: message}
}
func notFoundError(message string) *BorrowError {
	return &BorrowError{Kind: "isbn_not_found", message: message}
}
func holdInsteadError(message string) *BorrowError {
	return &BorrowError{Kind: "hold_instead", message: message}
}

// TODO(mchenryc): what should I be doing with this message? Should I just be
// logging outside of the business log e.g. when I handle the returned error
// of this func

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
func (i *Interactor) Borrow(ctx context.Context, isbn string, user model.User) (*model.Loan, *BorrowError) {
	logger := logging.From(ctx).With(
		zap.String("method", "borrow"),
		zap.String("user", user.Username),
		zap.String("isbn", isbn),
	)
	logger.Info("start")
	if !i.Authenticator.IsPatron(ctx, user) {
		logger.Info("unauthorized")
		return nil, unauthorizedError("user must be a patron to borrow a title")
	}

	titleCh := make(chan Result[*model.Title])
	countCh := make(chan Result[int])

	go func() {
		title, err := i.TitleRW.GetByIsbn(ctx, isbn)
		titleCh <- Result[*model.Title]{Data: title, Err: err}
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
		return nil, notFoundError(fmt.Sprintf("no title found with isbn %q", isbn))
	}

	countResult := <-countCh
	count, err := countResult.Unwrap()
	if err != nil {
		logger.Info("unable to load number of loans")
		return nil, internalError(fmt.Sprintf("unable to load number of loans for title with isbn %q", isbn))
	}

	// >= instead of == incase number of loans gets into a bad state
	if count >= title.Stock {
		logger.Info("title not available")
		return nil, holdInsteadError("title not available")
	}

	loan, err := i.LoanRW.Add(ctx, isbn, user)
	if err != nil {
		logger.Info("failed to add loan")
		return nil, internalError("failed to add loan")
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
