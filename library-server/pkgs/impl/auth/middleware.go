package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
)

type claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		c := claims{}
		keyFunc := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// TODO(mchenryc): get real secret
			return []byte("secret"), nil
		}

		// Verify Token
		_, err := jwt.ParseWithClaims(tokenString, &c, keyFunc)

		if err == nil {
			r = addUserContext(r, c)
			next.ServeHTTP(w, r)
		} else {
			// Unauthorized
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

type contextKey string

func addUserContext(r *http.Request, c claims) *http.Request {
	// TODO(mchenryc): add model.User to context
	key := contextKey("user")
	return r.WithContext(context.WithValue(r.Context(), key, c.User))
}
