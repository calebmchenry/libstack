package usecase

import (
	"context"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"
	"libstack/pkgs/util"

	"go.uber.org/zap"
)

func (i *Interactor) ListTitles(ctx context.Context, user model.User) ([]model.Title, *UseCaseError) {
	var emptyTitles []model.Title
	trace, _ := util.Uuid()
	logger := logging.From(ctx).With(
		zap.String("method", "addTitle"),
		zap.String("user", user.Username),
		zap.String("trace-id", trace),
	)
	logger.Info("start")
	if !user.IsLibrarian() && !user.IsPatron() {
		logger.Info("unauthorized")
		return emptyTitles, unauthorizedError("user must be a librarian or patron to add a title")
	}

	// TODO(mchenryc): implement cursor and paging
	titles, err := i.TitleRW.GetAll(ctx, user)
	if err != nil {
		logger.Info("unauthorized")
		return emptyTitles, unauthorizedError("user must be a librarian or patron to add a title")
	}
	return titles, nil
}
