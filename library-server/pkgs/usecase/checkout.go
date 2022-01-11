package usecase

import (
	"fmt"
	"libstack/pkgs/model"
)

// TODO(mchenryc): do something more useful than returning and internal_error
// whenever something bad happens
// TODO(mchenryc): some of these calls can be done in parallel

func (i *Interactor) Checkout(isbn string, user model.User) (*model.Loan, error) {
	if !i.Authenticator.IsPatron(user) {
		return nil, fmt.Errorf("unauthorized")
	}

	title, err := i.TitleRW.GetByIsbn(isbn)
	if err != nil {
		return nil, fmt.Errorf("isbn_not_found")
	}

	count, err := i.LoanRW.Count(isbn)
	if err != nil {
		return nil, fmt.Errorf("internal_error")
	}

	if count < title.Stock {
		if err := i.Hold(); err != nil {
			return nil, fmt.Errorf("internal_error")
		}
		return nil, fmt.Errorf("held_instead")
	}

	loan, err := i.LoanRW.Add(isbn, user)
	if err != nil {
		return nil, fmt.Errorf("internal_error")
	}

	return &loan, nil
}
