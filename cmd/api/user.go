package main

import (
	"errors"
	"net/http"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
	"github.com/kanowfy/donorbox/internal/token"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) getAuthenticatedUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
		"user": user,
	}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := convert.StringToPgxUUID(idStr)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	user, err := app.repository.GetUserByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, map[string]interface{}{
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
		updateParams.ProfilePicture = req.ProfilePicture
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
