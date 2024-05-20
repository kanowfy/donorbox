package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/token"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.status = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (app *application) enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Check if the request is a preflight request
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

			// Write a 200 OK status and return without further actions.
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{w, 200}
		start := time.Now()
		next.ServeHTTP(rec, r)
		duration := time.Since(start)

		slog.Info(
			fmt.Sprintf("HTTP %s %s responded with status %d in %vms", r.Method, r.URL.Path, rec.status, duration.Milliseconds()),
			slog.String("requestMethod", r.Method),
			slog.String("requestPath", r.URL.Path),
			slog.Int("statusCode", rec.status),
			slog.Int64("timeElapsedMS", duration.Milliseconds()),
		)
	})
}

func (app *application) requireUserAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract and verify token
		id, err := token.VerifyRequestToken(r)
		if err != nil {
			if errors.Is(err, token.ErrMissingToken) {
				app.authenticationRequiredResponse(w, r)
			} else {
				app.invalidAuthenticationTokenResponse(w, r)
			}
			return
		}

		user, err := app.repository.GetUserByID(r.Context(), convert.MustStringToPgxUUID(id))
		if err != nil {
			app.invalidCredentialsResponse(w, r)
			return
		}

		r = app.contextSetUser(r, &user)
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireEscrowAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := token.VerifyRequestToken(r)
		if err != nil {
			if errors.Is(err, token.ErrMissingToken) {
				app.authenticationRequiredResponse(w, r)
			} else {
				app.invalidAuthenticationTokenResponse(w, r)
			}
			return
		}

		user, err := app.repository.GetEscrowUserByID(r.Context(), convert.MustStringToPgxUUID(id))
		if err != nil {
			app.invalidCredentialsResponse(w, r)
			return
		}

		r = app.contextSetEscrowUser(r, &user)
		next.ServeHTTP(w, r)

	})
}
