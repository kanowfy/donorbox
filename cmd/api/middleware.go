package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/kanowfy/donorbox/internal/auth"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.status = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
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

func (app *application) requireAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract and verify token
		val := r.Header.Get("Authorization")
		if val == "" {
			app.authenticationRequiredResponse(w, r)
			return
		}

		parts := strings.Split(val, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		token := parts[1]

		id, err := auth.VerifyToken(token)
		if err != nil {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := app.repository.GetUserByID(r.Context(), mustStringToPgxUUID(id))
		if err != nil {
			app.invalidCredentialsResponse(w, r)
			return
		}

		r = app.contextSetUser(r, &user)
		next.ServeHTTP(w, r)
	})
}
