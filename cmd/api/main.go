package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/log"
)

type application struct {
	config config
	db     *db.Queries
}

func init() {
	logger := log.New(os.Stdout, slog.LevelDebug)
	slog.SetDefault(logger)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("error loading config: %v", err))
		os.Exit(1)
	}

	dbpool, err := openDB(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("error connecting to database: %v", err))
		os.Exit(1)
	}

	app := &application{
		config: cfg,
		db:     db.New(dbpool),
	}

	err = app.serve()
	if err != nil {
		slog.Error(err.Error())
	}
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, cfg.Dsn)
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
