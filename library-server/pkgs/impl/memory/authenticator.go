package memory

import (
	"context"
	"libstack/pkgs/model"
)

type Authenticator struct{}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) IsPatron(ctx context.Context, user model.User) bool {
	for _, v := range user.Roles {
		if v == "patron" {
			return true
		}
	}
	return false
}
