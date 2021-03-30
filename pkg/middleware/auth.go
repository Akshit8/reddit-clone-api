// Package middleware contains all http middlewares
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/pkg/token"
	"github.com/Akshit8/reddit-clone-api/pkg/user"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

// Auth decodes the JWT and packs the session into context
func Auth(tokenMaker token.Maker, userService user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if bearerToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			jwtToken := strings.Split(bearerToken, "Bearer ")[1]
			payload, err := tokenMaker.VerifyToken(jwtToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			// get the user from the database
			user, err := userService.GetUserByUsername(r.Context(), payload.Username)
			if err != nil {
				http.Error(w, "Inavlid token", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// FindUserFromContext finds the user from the context. REQUIRES Middleware to have run.
func FindUserFromContext(ctx context.Context) *entity.User {
	raw, _ := ctx.Value(userCtxKey).(*entity.User)
	return raw
}
