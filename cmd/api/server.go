package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.config.Host, app.config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("starting server", slog.String("addr", srv.Addr))

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
