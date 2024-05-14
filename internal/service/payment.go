package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kanowfy/donorbox/internal/db"
)

// TODO: create service interface & struct with pool and queries as fields when fully migrate to service layer

func AcceptBacking(ctx context.Context, dbpool *pgxpool.Pool, queries *db.Queries, arg db.CreateBackingParams) error {
	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := queries.WithTx(tx)

	/* process payment:
	- get escrow bank account information
	- transfer funds
	- if err, roll back
	*/

	backing, err := qtx.CreateBacking(ctx, arg)
	if err != nil {
		return err
	}

	if err = qtx.UpdateProjectFund(ctx, db.UpdateProjectFundParams{
		ID:            arg.ProjectID,
		BackingAmount: arg.Amount,
	}); err != nil {
		return err
	}

	_, err = qtx.CreateTransaction(ctx, db.CreateTransactionParams{
		BackingID:       backing.ID,
		TransactionType: db.TransactionTypeBacking,
		Amount:          arg.Amount,
		InitiatorID:     arg.BackerID,
		RecipientID:     arg.BackerID, // substitute with escrow id
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
