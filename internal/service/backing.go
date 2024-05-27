package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/models"
)

func (s *Service) AcceptBacking(ctx context.Context, projectID pgtype.UUID, request models.BackingRequest) error {
	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.repository.WithTx(tx)

	card, err := s.getCardToken(ctx, request.CardInformation)
	if err != nil {
		return err
	}

	escrow, err := qtx.GetEscrowUser(ctx)
	if err != nil {
		return err
	}

	backingParams := db.CreateBackingParams{
		ProjectID:     projectID,
		Amount:        convert.MustStringToInt64(request.Amount),
		WordOfSupport: request.WordOfSupport,
	}

	if request.UserID != nil {
		backingParams.BackerID = convert.MustStringToPgxUUID(*request.UserID)
	}

	backing, err := qtx.CreateBacking(ctx, backingParams)
	if err != nil {
		return err
	}

	if err = qtx.UpdateProjectFund(ctx, db.UpdateProjectFundParams{
		ID:            backing.ProjectID,
		BackingAmount: backing.Amount,
	}); err != nil {
		return err
	}

	_, err = qtx.CreateTransaction(ctx, db.CreateTransactionParams{
		ProjectID:       backing.ProjectID,
		TransactionType: db.TransactionTypeBacking,
		Amount:          backing.Amount,
		InitiatorCardID: card.ID,
		RecipientCardID: escrow.CardID,
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
