package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/token"
	"github.com/markbates/goth/gothic"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.repository.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// validate password
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password)); err != nil {
		app.invalidCredentialsResponse(w, r)
		return
	}

	token, err := token.GenerateToken(convert.PgxUUIDToString(user.ID), time.Hour*3*24)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// return token
	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) registerAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterAccountRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	args := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	user, err := app.repository.CreateUser(r.Context(), args)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	token, err := token.GenerateToken(convert.PgxUUIDToString(user.ID), time.Hour*3*24)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		payload := map[string]interface{}{
			"activationUrl": fmt.Sprintf("http://%s:%d/verify?token=%s", app.config.Host, 5173, token), // adjust url as needed
		}

		if err := app.mailer.Send(req.Email, "registration.tmpl", payload); err != nil {
			app.logError(r, err)
		}
	})

	if err = app.writeJSON(w, http.StatusAccepted, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) passwordForgotHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// check if email exist
	// if not send 201 Accepted, if yet, generate jwt to put in a link

	user, err := app.repository.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		if err = app.writeJSON(w, http.StatusAccepted, nil, nil); err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}

	token, err := token.GenerateToken(convert.PgxUUIDToString(user.ID), time.Minute*15)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		payload := map[string]interface{}{
			"firstName":        user.FirstName,
			"resetPasswordUrl": fmt.Sprintf("http://%s:%d/password/reset?token=%s", app.config.Host, 5173, token), // adjust url as needed
		}

		if err := app.mailer.Send(req.Email, "resetpassword.tmpl", payload); err != nil {
			app.logError(r, err)
		}
	})

	if err = app.writeJSON(w, http.StatusOK, nil, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ResetPasswordRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	id, err := token.VerifyToken(req.ResetToken)
	if err != nil {
		if errors.Is(err, token.ErrMissingToken) {
			app.badRequestResponse(w, r, err)
		} else {
			app.invalidAuthenticationTokenResponse(w, r)
		}
		return
	}

	user, err := app.repository.GetUserByID(r.Context(), convert.MustStringToPgxUUID(id))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	newHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	args := db.UpdateUserPasswordParams{
		ID:             user.ID,
		HashedPassword: string(newHashedPassword),
	}

	if err = app.repository.UpdateUserPassword(r.Context(), args); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "password reset successfully",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

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
			Email:          gothUser.Email,
			FirstName:      gothUser.FirstName,
			LastName:       gothUser.LastName,
			ProfilePicture: &gothUser.AvatarURL,
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
