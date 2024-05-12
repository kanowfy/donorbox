package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/token"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) getOneUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

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
			"activationUrl": fmt.Sprintf("%s:%d/verify?token=%s", app.config.Host, app.config.Port, token), // adjust url as needed
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

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := readString(r.URL.Query(), "token", "")
	if tokenString == "" {
		app.badRequestResponse(w, r, errors.New("missing token"))
		return
	}

	userID, err := token.VerifyToken(tokenString)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.repository.GetUserByID(r.Context(), convert.MustStringToPgxUUID(userID))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid token"))
		return
	}

	if user.Activated {
		app.badRequestResponse(w, r, errors.New("user has already been verified"))
		return
	}

	if err = app.repository.ActivateUser(r.Context(), user.ID); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "profile updated successfully",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateAccountRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	var updateParams db.UpdateUserByIDParams
	updateParams.ID = user.ID
	updateParams.Activated = user.Activated

	if req.Email != nil {
		updateParams.Email = *req.Email
	} else {
		updateParams.Email = user.Email
	}

	if req.FirstName != nil {
		updateParams.FirstName = *req.FirstName
	} else {
		updateParams.FirstName = user.FirstName
	}

	if req.LastName != nil {
		updateParams.LastName = *req.LastName
	} else {
		updateParams.LastName = user.LastName
	}

	if req.ProfilePicture != nil {
		pp := pgtype.Text{
			String: *req.ProfilePicture,
			Valid:  true,
		}

		updateParams.ProfilePicture = pp
	} else {
		updateParams.ProfilePicture = user.ProfilePicture
	}

	if err = app.repository.UpdateUserByID(r.Context(), updateParams); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "profile updated successfully",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ChangePasswordRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = app.validator.Struct(req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.OldPassword))
	if err != nil {
		app.badRequestResponse(w, r, err)
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
		"message": "password changed successfully",
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
