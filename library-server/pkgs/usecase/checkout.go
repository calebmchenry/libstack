package usecase

import (
	"fmt"
	"libstack/pkgs/model"
)

// TODO(mchenryc): do something more useful than returning and internal_error
// whenever something bad happens
// TODO(mchenryc): some of these calls can be done in parallel

// Checkout will create a loan of the specified title for the provided user.
//
// Use Case:
// (1) Only patrons can checkout titles
// (2) Only valid titles can be checkout
// (3) Only titles in stock can be checkedout
//
// Implemtation Details:
// (1) -> `"unauthorized"`
// (2) -> `"isbn_not_found"`
// (3) -> `"hold_instead"`
// (unexpected behavior) -> `"internal_error"`
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
		return nil, fmt.Errorf("hold_instead")
	}

	loan, err := i.LoanRW.Add(isbn, user)
	if err != nil {
		return nil, fmt.Errorf("internal_error")
	}

	return &loan, nil
}
