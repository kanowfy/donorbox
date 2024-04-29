package main

import (
	"net/http"
	"time"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/markbates/goth/gothic"
)

func (app *application) googleAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var user db.User

	user, err = app.repository.GetUserByEmail(r.Context(), user.Email)
	if err != nil {
		user, err = app.repository.CreateUser(r.Context(), db.CreateUserParams{
			Email:          gothUser.Email,
			FirstName:      gothUser.FirstName,
			LastName:       gothUser.LastName,
			HashedPassword: "",
		})

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	token, err := token.GenerateToken(convert.PgxUUIDToString(user.ID), time.Hour*3*24)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) startGoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (app *application) googleAuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
