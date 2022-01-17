package memory

import (
	"context"
	"fmt"
	"libstack/pkgs/model"
	"libstack/pkgs/util"

	"github.com/pkg/errors"
)

type LoanReadWriter struct {
	cache map[string]model.Loan
}

func NewLoanReadWriter() *LoanReadWriter {
	return &LoanReadWriter{cache: map[string]model.Loan{}}
}

func (l *LoanReadWriter) Add(ctx context.Context, isbn string, user model.User) (loan model.Loan, err error) {
	// TODO(mchenryc): logging
	// TODO(mchenryc): should validation happen somewhere else?
	if isbn == "" {
		err = errors.New("isbn required")
		return
	}
	id, err := util.Uuid()
	if err != nil {
		err = errors.Wrap(err, "failed to genereate unique Id")
		return
	}
	loan = model.Loan{Id: id, Isbn: isbn, Username: user.Username}
	l.cache[loan.Id] = loan
	return loan, nil
}

func (l *LoanReadWriter) Count(ctx context.Context, isbn string) (int, error) {
	// TODO(mchenryc): logging
	// TODO(mchenryc): should validation happen somewhere else?
	if isbn == "" {
		return 0, fmt.Errorf("isbn required")
	}
	count := 0
	for _, v := range l.cache {
		if v.Isbn == isbn {
			count++
		}
	}
	return count, nil
}
