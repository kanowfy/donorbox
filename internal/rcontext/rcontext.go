package rcontext

import (
	"context"
	"net/http"

	"github.com/kanowfy/donorbox/internal/model"
)

type contextKey string

const (
	userContextKey   = contextKey("user")
	escrowContextKey = contextKey("escrow")
)

// SetUser attaches current user to the request context.
func SetUser(r *http.Request, user *model.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func GetUser(r *http.Request) *model.User {
	user, ok := r.Context().Value(userContextKey).(*model.User)
	if !ok {
		panic("missing user value in request")
	}

	return user
}

// SetEscrowUser attaches current escrow user to the request context.
func SetEscrowUser(r *http.Request, escrow *model.EscrowUser) *http.Request {
	ctx := context.WithValue(r.Context(), escrowContextKey, escrow)
	return r.WithContext(ctx)
}

func GetEscrowUser(r *http.Request) *model.EscrowUser {
	user, ok := r.Context().Value(escrowContextKey).(*model.EscrowUser)
	if !ok {
		panic("missing escrow user value in request")
	}

	return user
}
