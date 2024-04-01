package main

import (
	"log/slog"
	"os"

	"github.com/kanowfy/donorbox/internal/log"
)

type application struct {
	config config
}

func init() {
	logger := log.New(os.Stdout, slog.LevelDebug)
	slog.SetDefault(logger)
}

func main() {
	cfg := config{
		host: "localhost",
		port: 4000,
	}

	app := &application{
		config: cfg,
	}

	err := app.serve()
	if err != nil {
		slog.Error(err.Error())
	}
}
