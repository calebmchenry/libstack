package usecase

import "libstack/pkgs/model"

type Interactor struct {
	Authenticator Authenticator
	TitleRW       TitleReadWriter
	LoanRW        LoanReadWriter
}

type Authenticator interface {
	IsLibrarian(user model.User) bool
	IsPatron(user model.User) bool
}

type TitleReadWriter interface {
	GetByIsbn(isbn string) (model.Title, error)
	Update(model.Title) (model.Title, error)
}

type LoanReadWriter interface {
	Add(isbn string, user model.User) (model.Loan, error)
	Count(isbn string) (int, error)
}
