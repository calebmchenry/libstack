package usecase

import (
	"context"
	"libstack/pkgs/model"
)

type Interactor struct {
	Authenticator Authenticator
	TitleRW       TitleReadWriter
	LoanRW        LoanReadWriter
}

type Authenticator interface {
	IsPatron(ctx context.Context, user model.User) bool
}

type TitleReadWriter interface {
	GetByIsbn(ctx context.Context, isbn string) (model.Title, error)
	Add(ctx context.Context, title model.Title) (model.Title, error)
}

type LoanReadWriter interface {
	Add(ctx context.Context, isbn string, user model.User) (model.Loan, error)
	Count(ctx context.Context, isbn string) (int, error)
}
