package main

import (
	"context"
	"net/http"

	"github.com/kanowfy/donorbox/internal/db"
)

type contextKey string

const (
	userContextKey   = contextKey("user")
	escrowContextKey = contextKey("escrow")
)

func (app *application) contextSetUser(r *http.Request, user *db.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *db.User {
	user, ok := r.Context().Value(userContextKey).(*db.User)
	if !ok {
		panic("missing user value in request")
	}

	return user
}

func (app *application) contextSetEscrowUser(r *http.Request, escrow *db.EscrowUser) *http.Request {
	ctx := context.WithValue(r.Context(), escrowContextKey, escrow)
	return r.WithContext(ctx)
}

func (app *application) contextGetEscrowUser(r *http.Request) *db.EscrowUser {
	user, ok := r.Context().Value(escrowContextKey).(*db.EscrowUser)
	if !ok {
		panic("missing escrow user value in request")
	}

	return user
}
