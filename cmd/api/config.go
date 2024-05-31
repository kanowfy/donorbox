package main

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"4000"`

	DBDsn          string `env:"DB_DSN"`
	DBUser         string `env:"DB_USER"`
	DBPassword     string `env:"DB_PASSWORD"`
	DBHost         string `env:"DB_HOST"`
	DBPort         int    `env:"DB_PORT" envDefault:"5432"`
	DBName         string `env:"DB_NAME"`
	DBSslMode      string `env:"DB_SSLMODE" envDefault:"disable"`
	DBMaxOpenConns int    `env:"DB_MAXOPENCONNS" envDefault:"25"`
	DBMaxIdleConns int    `env:"DB_MAXIDLECONNS" envDefault:"25"`
	DBMaxIdleTime  string `env:"DB_MAXIDLETIME" envDefault:"15m"`

	ClientPort int `env:"CLIENT_PORT" envDefault:"4001"`

	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     int    `env:"SMTP_PORT" envDefault:"25"`
	SmtpUsername string `env:"SMTP_USERNAME"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
	SmtpSender   string `env:"SMTP_SENDER" envDefault:"DonorBox <no-reply@donorbox.kanowfy.com>"`

	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`

	CloudinaryAPIUrl string `env:"CLOUDINARY_API_URL"`
}

func loadConfig() (config, error) {
	godotenv.Load(".env")
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
