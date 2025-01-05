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
	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/publish"
	"github.com/kanowfy/donorbox/internal/router"
	"github.com/kanowfy/donorbox/internal/service"
)

func (app *application) run() error {
	ctx, stop := signal.NotifyContext(app.ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// initialize repository, usually repository is also broken up to correspond to each domain model
	// but here we group them together as provided by sqlc
	repository := db.New(app.dbpool)
	notifChan := make(chan model.Notification)
	publisher := publish.New(notifChan)

	// initialize service
	authService := service.NewAuth(repository, app.mailer)
	auditService := service.NewAuditTrail(repository)
	userService := service.NewUser(repository, auditService)
	escrowService := service.NewEscrow(repository, app.mailer, publisher, auditService)
	backingService := service.NewBacking(repository, auditService)
	projectService := service.NewProject(repository, backingService, userService, auditService)
	notificationService := service.NewNotification(repository)
	ragService := service.NewRag(app.weaviateClient, app.genModel, app.embedModel)

	// initialize handlers
	authHandler := handler.NewAuth(authService, app.validator, app.cfg)
	auditHandler := handler.NewAuditTrail(auditService)
	userHandler := handler.NewUser(userService, app.validator, app.cfg, app.cfg.DropboxAccessToken)
	escrowHandler := handler.NewEscrow(escrowService, app.validator)
	backingHandler := handler.NewBacking(backingService, app.validator)
	projectHandler := handler.NewProject(projectService, app.validator)
	imageUploadHandler := handler.NewImageUploader(app.cfg)
	notifcationHandler := handler.NewNotification(notificationService, notifChan, ctx)
	ragHandler := handler.NewRag(ragService)

	// initialize auth middleware
	authMiddleware := middleware.NewAuth(userService, escrowService)

	handlers := handler.Handlers{
		Auth:          authHandler,
		Backing:       backingHandler,
		Escrow:        escrowHandler,
		Project:       projectHandler,
		User:          userHandler,
		ImageUploader: imageUploadHandler,
		Notification:  notifcationHandler,
		Rag:           ragHandler,
		AuditTrail:    auditHandler,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", app.cfg.Port),
		Handler:      router.Setup(handlers, authMiddleware, app.cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	cronJobs := cron.New(projectService)
	cronJobs.RunDaily()

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
