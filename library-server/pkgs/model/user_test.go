package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPassword(t *testing.T) {
	t.Run("does not allow empty password", func(t *testing.T) {
		user := User{}
		password := ""
		err := user.setPassword(password)

		assert.NotNil(t, err)
	})

	t.Run("hashes provided password", func(t *testing.T) {
		user := User{}
		password := "password123"
		user.setPassword(password)

		assert.NotEqual(t, password, user.PasswordHash)
	})

}

func TestCheckPassword(t *testing.T) {
	t.Run("returns no error when password matches", func(t *testing.T) {
		user := User{}
		password := "password123"
		user.setPassword(password)

		err := user.checkPassword(password)
		assert.Nil(t, err)
	})

	t.Run("returns error when password does not matches", func(t *testing.T) {
		user := User{}
		password := "password123"
		otherPassword := "123password"
		user.setPassword(password)

		err := user.checkPassword(otherPassword)
		assert.NotNil(t, err)
	})
}
