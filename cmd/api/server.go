package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		slog.Info(fmt.Sprintf("starting server at %s", srv.Addr))
		err := srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("server error: %v", err))
		}
	}()

	<-ctx.Done()
	slog.Info("caught interruption signal, shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}
