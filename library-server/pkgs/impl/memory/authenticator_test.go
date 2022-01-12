package memory_test

import (
	"context"
	"libstack/pkgs/impl/memory"
	"libstack/pkgs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_IsPatron(t *testing.T) {
	t.Run("returns true if roles contains 'patron'", func(t *testing.T) {
		user := model.User{Roles: []string{"patron"}}
		a := memory.NewAuthenticator()
		assert.True(t, a.IsPatron(context.Background(), user))
	})
	t.Run("returns false if roles does not contain 'patron'", func(t *testing.T) {
		user := model.User{Roles: []string{"foo"}}
		a := memory.NewAuthenticator()
		assert.False(t, a.IsPatron(context.Background(), user))
	})

}
