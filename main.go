package main

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/log"
	"github.com/kanowfy/donorbox/internal/mail"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type application struct {
	cfg       config.Config
	dbpool    *pgxpool.Pool
	validator *validator.Validate
	mailer    mail.Mailer
	wg        sync.WaitGroup
}

func init() {
	logger := log.New(os.Stdout, slog.LevelDebug)
	slog.SetDefault(logger)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("error loading config: %v", err))
		os.Exit(1)
	}

	dbpool, err := openDB(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("error connecting to database: %v", err))
		os.Exit(1)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDBFromPool(dbpool)

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	goth.UseProviders(
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleClientSecret,
			fmt.Sprintf("http://%s:%d/api/v1/users/auth/google/callback", cfg.Host, cfg.Port),
		),
	)

	app := &application{
		cfg:       cfg,
		dbpool:    dbpool,
		validator: validator.New(validator.WithRequiredStructEnabled()),
		mailer:    mail.New(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpUsername, cfg.SmtpPassword, cfg.SmtpSender),
	}

	err = app.run()
	if err != nil {
		slog.Error(err.Error())
	}
}

func openDB(cfg config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_idle_time=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSslMode, cfg.DBMaxIdleConns, cfg.DBMaxIdleConns, cfg.DBMaxIdleTime))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = dbpool.Ping(ctx); err != nil {
		return nil, err
	}

	return dbpool, nil
}
