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

	"github.com/kanowfy/donorbox/internal/cron"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/handler"
	"github.com/kanowfy/donorbox/internal/middleware"
	"github.com/kanowfy/donorbox/internal/router"
	"github.com/kanowfy/donorbox/internal/service"
)

func (app *application) run() error {
	// initialize repository, usually repository is also broken up to correspond to each domain model
	// but here we group them together as provided by sqlc
	repository := db.New(app.dbpool)

	// initialize service
	authService := service.NewAuth(repository, app.mailer)
	userService := service.NewUser(repository)
	escrowService := service.NewEscrow(repository)
	backingService := service.NewBacking(repository)
	projectService := service.NewProject(repository, backingService, userService)

	// initialize handlers
	authHandler := handler.NewAuth(authService, app.validator, app.cfg)
	userHandler := handler.NewUser(userService, app.validator, app.cfg, app.cfg.DropboxAccessToken)
	escrowHandler := handler.NewEscrow(escrowService, app.validator)
	backingHandler := handler.NewBacking(backingService, app.validator)
	projectHandler := handler.NewProject(projectService, app.validator)
	imageUploadHandler := handler.NewImageUploader(app.cfg)

	// initialize auth middleware
	authMiddleware := middleware.NewAuth(userService, escrowService)

	handlers := handler.Handlers{
		Auth:          authHandler,
		Backing:       backingHandler,
		Escrow:        escrowHandler,
		Project:       projectHandler,
		User:          userHandler,
		ImageUploader: imageUploadHandler,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", app.cfg.Port),
		Handler:      router.Setup(handlers, authMiddleware, app.cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	cronJobs := cron.New(projectService)
	cronJobs.Start()

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
	defer func() {
		slog.Info("completing background tasks")
		app.wg.Wait()
	}()
	slog.Info("caught interruption signal, shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}
