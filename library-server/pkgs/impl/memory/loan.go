package memory

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"libstack/pkgs/model"
)

type LoanReadWriter struct {
	cache map[string]model.Loan
}

func NewLoanReadWriter() *LoanReadWriter {
	return &LoanReadWriter{cache: map[string]model.Loan{}}
}

// TODO(mchenryc): move uuid to a better location

// NewUUID generates a random UUID according to RFC 4122
func uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func (l *LoanReadWriter) Add(ctx context.Context, isbn string, user model.User) (*model.Loan, error) {
	// TODO(mchenryc): logging
	// TODO(mchenryc): should validation happen somewhere else?
	// TODO(mchenryc): better errors
	if isbn == "" {
		return nil, fmt.Errorf("isbn required")
	}
	id, err := uuid()
	if err != nil {
		return nil, fmt.Errorf("failed to genereate unique Id")
	}
	loan := model.Loan{Id: id, Isbn: isbn, Username: user.Username}
	l.cache[loan.Id] = loan
	return &loan, nil
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
