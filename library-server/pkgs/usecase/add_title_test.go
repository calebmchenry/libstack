package usecase_test

import (
	"context"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTitle(t *testing.T) {
	t.Run("Adds title when librarian adds a valid title", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"librarian"}}

		// TODO(mchenryc): assert title to be added is not in listing

		title := model.Title{Isbn: isbn, Stock: 1}
		title, err := i.AddTitle(ctx, title, user)

		// TODO(mchenryc): assert added title get be retrieve from listing

		assert.NotNil(t, title)
		assert.Nil(t, err)
	})
	t.Run("does not permit non librarians to add a title", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := "1234"
		user := model.User{Roles: []string{"patron"}}

		title := model.Title{Isbn: isbn, Stock: 1}
		title, err := i.AddTitle(ctx, title, user)

		assert.NotNil(t, err)
		assert.Equal(t, "unauthorized", err.Kind)
	})
	t.Run("returns invalid isbn if title does not have an isbn", func(t *testing.T) {
		i := setupInteractor()

		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)

		isbn := ""
		user := model.User{Roles: []string{"librarian"}}

		title := model.Title{Isbn: isbn, Stock: 1}
		_, err := i.AddTitle(ctx, title, user)
		assert.NotNil(t, err)
	})
	t.Run("returns internal error when title failed to be persisted", func(t *testing.T) {})
}
