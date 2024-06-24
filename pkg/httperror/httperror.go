package httperror

import (
	"log/slog"
	"net/http"

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
