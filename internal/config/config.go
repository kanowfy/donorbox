package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
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

	ClientUrl    string `env:"CLIENT_URL" envDefault:"http://localhost:4001"`
	DashboardUrl string `env:"DASHBOARD_URL" envDefault:"http://localhost:4002"`

	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     int    `env:"SMTP_PORT" envDefault:"25"`
	SmtpUsername string `env:"SMTP_USERNAME"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
	SmtpSender   string `env:"SMTP_SENDER" envDefault:"DonorBox <no-reply@donorbox.kanowfy.com>"`

	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`

	CloudinaryAPIUrl string `env:"CLOUDINARY_API_URL"`

	StripeSecretKey string `env:"STRIPE_SK"`

	DropboxAccessToken  string `env:"DROPBOX_ACCESS_TOKEN"`
	DropboxRefreshToken string `env:"DROPBOX_REFRESH_TOKEN"`
	DropboxAppKey       string `env:"DROPBOX_APP_KEY"`
	DropboxAppSecret    string `env:"DROPBOX_APP_SECRET"`
}

func Load() (Config, error) {
	godotenv.Load(".env")
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
