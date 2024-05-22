package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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

	fmt.Printf("%+v\n", gothUser)

	var user db.User

	user, err = app.repository.GetUserByEmail(r.Context(), gothUser.Email)
	if err != nil {
		params := db.CreateSocialLoginUserParams{
			Email:     gothUser.Email,
			FirstName: gothUser.FirstName,
			LastName:  gothUser.LastName,
			ProfilePicture: pgtype.Text{
				String: gothUser.AvatarURL,
				Valid:  true,
			},
		}

		if params.FirstName == "" {
			params.FirstName = "Anonymous"
		}

		user, err = app.repository.CreateSocialLoginUser(r.Context(), params)

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

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 3 * 24),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.Redirect(w, r, "http://localhost:5173/login/google", http.StatusFound)
}

func (app *application) getAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token": cookie.Value,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) startGoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}

func (app *application) googleAuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
