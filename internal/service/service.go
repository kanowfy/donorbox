package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanowfy/donorbox/internal/db"
)

type Service struct {
	repository *db.Queries
	dbPool     *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) *Service {
	return &Service{
		repository: db.New(dbpool),
		dbPool:     dbpool,
	}
}
