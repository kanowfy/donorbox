package main

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         int    `env:"PORT" envDefault:"4000"`
	Dsn          string `env:"DB_DSN"`
	MaxOpenConns int    `env:"DB_MAXOPENCONNS" envDefault:"25"`
	MaxIdleConns int    `env:"DB_MAXIDLECONNS" envDefault:"25"`
	MaxIdleTime  string `env:"DB_MAXIDLETIME" envDefault:"15m"`
}

func loadConfig() (config, error) {
	godotenv.Load(".env")
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
