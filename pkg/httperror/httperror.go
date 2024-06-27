package httperror

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/pkg/json"
)

func LogError(r *http.Request, err error) {
	slog.Error(
		err.Error(),
		slog.String("request_method", r.Method),
		slog.String("request_url", r.URL.String()),
	)
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	m := map[string]interface{}{
		"error": message,
	}

	err := json.WriteJSON(w, status, m, nil)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, err)

	errorResponse(w, r, http.StatusInternalServerError, "could not process the request")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusNotFound, "could not find the requested resource")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusUnauthorized, "invalid credentials")
}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	errorResponse(w, r, http.StatusUnauthorized, "invalid authentication token")
}

func AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusUnauthorized, "you must be authenticated to access this resource")
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("err has to be validation error")
	}

	type inputError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	errors := []inputError{}
	for _, e := range errs {
		var inputErr inputError
		inputErr.Field = e.Field()
		switch e.Tag() {
		case "required":
			inputErr.Message = fmt.Sprintf("missing required field")
		case "email":
			inputErr.Message = "invalid email"
		case "credit_card":
			inputErr.Message = "invalid credit card value"
		case "uuid4":
			inputErr.Message = "invalid uuid"
		case "http_url":
			inputErr.Message = "invalid url"
		default:
			inputErr.Message = fmt.Sprintf("validation failed on '%s' tag", e.Tag())
		}

		errors = append(errors, inputErr)
	}

	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
