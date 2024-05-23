package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/db"
)

func (s *Service) Payout(ctx context.Context, projectID pgtype.UUID, escrow *db.EscrowUser) error {
	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.repository.WithTx(tx)

	/* flow for payout
	projectID -> set all backing releated to project to released, set project status to completed_payout
	-> create a new transaction
	*/

	if err = qtx.UpdateProjectBackingStatus(ctx, db.UpdateProjectBackingStatusParams{
		ProjectID: projectID,
		Status:    db.BackingStatusReleased,
	}); err != nil {
		return err
	}

	if err = qtx.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
		ID:     projectID,
		Status: db.ProjectStatusCompletedPayout,
	}); err != nil {
		return err
	}

	project, err := qtx.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}

	_, err = qtx.CreateTransaction(ctx, db.CreateTransactionParams{
		ProjectID:       projectID,
		TransactionType: db.TransactionTypePayout,
		Amount:          project.CurrentAmount,
		InitiatorCardID: escrow.CardID,
		RecipientCardID: project.CardID,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *Service) Refund(ctx context.Context, projectID pgtype.UUID, escrow *db.EscrowUser) error {
	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.repository.WithTx(tx)

	/* flow for refund
	projectID -> set all backing releated to project to refuned, set project status to completed_refund
	-> create a new transaction
	*/

	if err = qtx.UpdateProjectBackingStatus(ctx, db.UpdateProjectBackingStatusParams{
		ProjectID: projectID,
		Status:    db.BackingStatusRefunded,
	}); err != nil {
		return err
	}

	if err = qtx.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
		ID:     projectID,
		Status: db.ProjectStatusCompletedRefund,
	}); err != nil {
		return err
	}

	transactions, err := qtx.GetBackingTransactionsForProject(ctx, projectID)
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		_, err = qtx.CreateTransaction(ctx, db.CreateTransactionParams{
			ProjectID:       projectID,
			TransactionType: db.TransactionTypeRefund,
			Amount:          transaction.Amount,
			InitiatorCardID: escrow.CardID,
			RecipientCardID: transaction.InitiatorCardID,
		})

		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}
