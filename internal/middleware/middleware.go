package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.status = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func EnableCors(trustedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//trustedOrigins := []string{fmt.Sprintf("http://%s:%d", app.config.Host, app.config.ClientPort), fmt.Sprintf("http://%s:%d", app.config.Host, app.config.DashboardPort)}
		w.Header().Set("Vary", "Origin")
		origin := r.Header.Get("Origin")

		if origin != "" {
			for i := range trustedOrigins {
				if origin == trustedOrigins[i] {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Credentials", "true")

					// Check if the request is a preflight request
					if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
						w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
						w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

						// Write a 200 OK status and return without further actions.
						w.WriteHeader(http.StatusOK)
						return
					}
					break

				}
			}

		}

		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
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
