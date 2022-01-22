package usecase_test

import (
	"context"
	"libstack/pkgs/impl/logging"
	"libstack/pkgs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTitles(t *testing.T) {
	t.Run("only allows patrons and librarians to view titles", func(t *testing.T) {
		i := setupInteractor()
		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)
		nobody := model.User{Roles: []string{"foo"}}

		_, err := i.ListTitles(ctx, nobody)
		assert.NotNil(t, err)
		assert.Equal(t, "unauthorized", err.Kind)

		patron := model.User{Roles: []string{"patron"}}
		_, err = i.ListTitles(ctx, patron)
		assert.Nil(t, err)

		librarian := model.User{Roles: []string{"librarian"}}
		_, err = i.ListTitles(ctx, librarian)
		assert.Nil(t, err)
	})

	t.Run("returns all titles", func(t *testing.T) {
		i := setupInteractor()
		logger, _ := logging.BufferLogger()
		ctx := logging.NewContext(context.Background(), logger)
		user := model.User{Roles: []string{"librarian"}}
		titles, err := i.ListTitles(ctx, user)
		assert.Nil(t, err)
		assert.Len(t, titles, 0)

		title := model.Title{Isbn: "1234"}
		i.AddTitle(ctx, title, user)

		titles, err = i.ListTitles(ctx, user)
		assert.Nil(t, err)
		assert.Len(t, titles, 1)
	})
}
