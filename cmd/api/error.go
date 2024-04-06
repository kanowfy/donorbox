package main

import (
	"log/slog"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	slog.Error(
		err.Error(),
		slog.String("request_method", r.Method),
		slog.String("request_url", r.URL.String()),
	)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	m := map[string]interface{}{
		"error": message,
	}

	err := app.writeJSON(w, status, m, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	app.errorResponse(w, r, http.StatusInternalServerError, "could not process the request")
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, "could not find the requested resource")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusNotFound, err.Error())
}
