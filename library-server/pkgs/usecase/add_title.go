package usecase

import (
	"context"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"
	"libstack/pkgs/util"

	"go.uber.org/zap"
)

func (i *Interactor) AddTitle(ctx context.Context, title model.Title, user model.User) (model.Title, *UseCaseError) {
	var emptyTitle model.Title

	trace, _ := util.Uuid()
	logger := logging.From(ctx).With(
		zap.String("method", "addTitle"),
		zap.String("user", user.Username),
		zap.String("isbn", title.Isbn),
		zap.String("trace-id", trace),
	)
	logger.Info("start")
	if !user.IsLibrarian() {
		logger.Info("unauthorized")
		return emptyTitle, unauthorizedError("user must be a librarian to add a title")
	}
	if title.Isbn == "" {
		logger.Info("invalid isbn, isbn is required but received an empty string ")
		return emptyTitle, notFoundError("isbn is required")
	}
	title, err := i.TitleRW.Add(ctx, title)
	if err != nil {
		logger.Info("failed to add loan")
		return emptyTitle, internalError("failed to add title", err)
	}

	return title, nil
}
