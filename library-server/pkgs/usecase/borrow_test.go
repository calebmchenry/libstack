package usecase_test

import (
	"context"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/impl/memory"
	"libstack/pkgs/model"
	"libstack/pkgs/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupInteractor() usecase.Interactor {
	loanRW := memory.NewLoanReadWriter()
	titleRW := memory.NewTitleReadWriter()
	authenticator := memory.NewAuthenticator()
	return usecase.Interactor{
		Authenticator: authenticator,
		LoanRW:        loanRW,
		TitleRW:       titleRW,
	}
}

func TestBorrow(t *testing.T) {
	t.Run("creates loan when patron borrows an available book", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"patron"}}

		title := model.Title{Isbn: isbn, Stock: 1}
		i.TitleRW.Add(ctx, title)

		initialCount, err := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 0, initialCount)
		assert.Nil(t, err)

		loan, err := i.Borrow(ctx, isbn, user)
		assert.NotNil(t, loan)
		assert.Nil(t, err)

		finalCount, err := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 1, finalCount)
		assert.Nil(t, err)
	})
	t.Run("does not permit non patrons to borrow a title", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"foo"}}

		title := model.Title{Isbn: isbn, Stock: 1}
		i.TitleRW.Add(ctx, title)

		initialCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 0, initialCount)

		loan, err := i.Borrow(ctx, isbn, user)
		assert.Nil(t, loan)
		assert.NotNil(t, err)
		assert.Contains(t, "unauthorized", err.Kind)

		finalCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 0, finalCount)
	})
	t.Run("returns isbn_not_found if no title exists for provided isbn", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"patron"}}

		initialCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 0, initialCount)

		loan, err := i.Borrow(ctx, isbn, user)
		assert.Nil(t, loan)
		assert.NotNil(t, err)
		assert.Contains(t, "isbn_not_found", err.Kind)

		finalCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 0, finalCount)
	})
	t.Run("handles failing to load loan count", func(t *testing.T) {

	})
	t.Run("prevents borrow if title is not availabe", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"patron"}}
		title := model.Title{Isbn: isbn, Stock: 1}
		i.TitleRW.Add(ctx, title)

		i.LoanRW.Add(ctx, isbn, user)

		initialCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 1, initialCount)

		loan, err := i.Borrow(ctx, isbn, user)
		assert.Nil(t, loan)
		assert.NotNil(t, err)
		assert.Contains(t, "hold_instead", err.Kind)

		finalCount, _ := i.LoanRW.Count(ctx, isbn)
		assert.Equal(t, 1, finalCount)
	})
	t.Run("handles failing to add loan", func(t *testing.T) {

	})
}
