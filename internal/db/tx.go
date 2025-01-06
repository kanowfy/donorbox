package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TX represents a database transaction
type TX interface {
	Commit(context.Context) error
	Rollback(context.Context) error
}

// BeginTXer represents a transaction initiator, which return a valid transaction and repository interface
type BeginTXer interface {
	BeginTX(ctx context.Context, opt pgx.TxOptions) (Querier, TX, error)
}

// BeginTX creates resources to run a database transaction
func (q *Queries) BeginTX(ctx context.Context, opt pgx.TxOptions) (Querier, TX, error) {
	txer, ok := q.db.(*pgxpool.Pool)
	if !ok {
		return nil, nil, errors.New("db is not a pgx connection pool")
	}

	tx, err := txer.BeginTx(ctx, opt)
	if err != nil {
		return nil, nil, err
	}

	return q.WithTx(tx), tx, nil
}
