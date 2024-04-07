package main

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
