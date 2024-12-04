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
	"github.com/google/generative-ai-go/genai"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/log"
	"github.com/kanowfy/donorbox/internal/mail"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/pressly/goose/v3"
	"github.com/stripe/stripe-go/v81"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"google.golang.org/api/option"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type application struct {
	ctx       context.Context
	cfg       config.Config
	dbpool    *pgxpool.Pool
	validator *validator.Validate
	mailer    mail.Mailer
	wg        sync.WaitGroup

	weaviateClient *weaviate.Client
	embedModel     *genai.EmbeddingModel
	genModel       *genai.GenerativeModel
}

func init() {
	logger := log.New(os.Stdout, slog.LevelDebug)
	slog.SetDefault(logger)
}

func main() {
	ctx := context.Background()

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

	stripe.Key = cfg.StripeSecretKey

	weaviateClient, err := initiateWeaviate(ctx, cfg)
	if err != nil {
		panic(err)
	}

	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiApiKey))
	if err != nil {
		panic(err)
	}

	app := &application{
		ctx:            ctx,
		cfg:            cfg,
		dbpool:         dbpool,
		validator:      validator.New(validator.WithRequiredStructEnabled()),
		mailer:         mail.New(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpUsername, cfg.SmtpPassword, cfg.SmtpSender),
		weaviateClient: weaviateClient,
		embedModel:     genaiClient.EmbeddingModel("text-embedding-004"),
		genModel:       genaiClient.GenerativeModel("gemini-1.5-flash"),
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

func initiateWeaviate(ctx context.Context, cfg config.Config) (*weaviate.Client, error) {
	client, err := weaviate.NewClient(weaviate.Config{
		Host:   "localhost:" + cfg.WeaviatePort,
		Scheme: "http",
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to weaviate: %w", err)
	}

	cls := &models.Class{
		Class:      "Document",
		Vectorizer: "none",
	}

	exists, err := client.Schema().ClassExistenceChecker().WithClassName(cls.Class).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error from weaviate: %w", err)
	}

	if !exists {
		if err = client.Schema().ClassCreator().WithClass(cls).Do(ctx); err != nil {
			return nil, fmt.Errorf("error from weaviate: %w", err)
		}
	}

	return client, nil
}
