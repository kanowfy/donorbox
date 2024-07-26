package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
)

type Backing interface {
	AcceptBacking(ctx context.Context, projectID uuid.UUID, request dto.BackingRequest) error
	GetBackingsForProject(ctx context.Context, projectID uuid.UUID) ([]model.Backing, error)
	GetProjectBackingAggregation(ctx context.Context, projectID uuid.UUID) (*BackingAggregation, error)
}

type backing struct {
	repository  db.Querier
	cardService Card
}

func NewBacking(repository db.Querier, cardService Card) Backing {
	return &backing{
		repository,
		cardService,
	}
}

func (b *backing) AcceptBacking(ctx context.Context, projectID uuid.UUID, request dto.BackingRequest) error {
	queries := b.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	card, err := b.cardService.RequestCardToken(ctx, request.CardInformation)
	if err != nil {
		return err
	}

	escrow, err := q.GetEscrowUser(ctx)
	if err != nil {
		return err
	}

	backingParams := db.CreateBackingParams{
		ProjectID:     projectID,
		Amount:        convert.MustStringToInt64(request.Amount),
		WordOfSupport: request.WordOfSupport,
	}

	if request.UserID != nil {
		backingParams.BackerID = uuid.MustParse(*request.UserID)
	}

	backing, err := q.CreateBacking(ctx, backingParams)
	if err != nil {
		return err
	}

	if err = q.UpdateProjectFund(ctx, db.UpdateProjectFundParams{
		ID:            backing.ProjectID,
		BackingAmount: backing.Amount,
	}); err != nil {
		return err
	}

	_, err = q.CreateTransaction(ctx, db.CreateTransactionParams{
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

func (b *backing) GetBackingsForProject(ctx context.Context, projectID uuid.UUID) ([]model.Backing, error) {
	_, err := b.repository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	dbBackings, err := b.repository.GetBackingsForProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	var backings []model.Backing
	for _, b := range dbBackings {
		backings = append(backings, model.Backing{
			ID:              b.ID,
			ProjectID:       b.ProjectID,
			Amount:          b.Amount,
			WordOfSupport:   b.WordOfSupport,
			Status:          convertBackingStatus(b.Status),
			CreatedAt:       b.CreatedAt,
			BackerFirstName: b.FirstName,
			BackerLastName:  b.LastName,
			ProfilePicture:  b.ProfilePicture,
		})
	}

	return backings, nil
}

type BackingAggregation struct {
	MostAmountBacking model.Backing
	FirstBacking      model.Backing
	RecentBacking     model.Backing
	TotalBacking      int64
}

func (b *backing) GetProjectBackingAggregation(ctx context.Context, projectID uuid.UUID) (*BackingAggregation, error) {
	mostBacking, err := b.repository.GetMostBackingDonor(ctx, projectID)
	if err != nil {
		return nil, err
	}

	firstBacking, err := b.repository.GetFirstBackingDonor(ctx, projectID)
	if err != nil {
		return nil, err
	}

	recentBacking, err := b.repository.GetMostRecentBackingDonor(ctx, projectID)
	if err != nil {
		return nil, err
	}

	count, err := b.repository.GetBackingCountForProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	backingAgg := &BackingAggregation{
		MostAmountBacking: model.Backing{
			ID:              mostBacking.ID,
			ProjectID:       mostBacking.ProjectID,
			Amount:          mostBacking.Amount,
			WordOfSupport:   mostBacking.WordOfSupport,
			Status:          convertBackingStatus(mostBacking.Status),
			CreatedAt:       mostBacking.CreatedAt,
			BackerFirstName: mostBacking.FirstName,
			BackerLastName:  mostBacking.LastName,
			ProfilePicture:  mostBacking.ProfilePicture,
		},
		FirstBacking: model.Backing{
			ID:              firstBacking.ID,
			ProjectID:       firstBacking.ProjectID,
			Amount:          firstBacking.Amount,
			WordOfSupport:   firstBacking.WordOfSupport,
			Status:          convertBackingStatus(firstBacking.Status),
			CreatedAt:       firstBacking.CreatedAt,
			BackerFirstName: firstBacking.FirstName,
			BackerLastName:  firstBacking.LastName,
			ProfilePicture:  firstBacking.ProfilePicture,
		},
		RecentBacking: model.Backing{
			ID:              recentBacking.ID,
			ProjectID:       recentBacking.ProjectID,
			Amount:          recentBacking.Amount,
			WordOfSupport:   recentBacking.WordOfSupport,
			Status:          convertBackingStatus(recentBacking.Status),
			CreatedAt:       recentBacking.CreatedAt,
			BackerFirstName: recentBacking.FirstName,
			BackerLastName:  recentBacking.LastName,
			ProfilePicture:  recentBacking.ProfilePicture,
		},
		TotalBacking: count,
	}

	return backingAgg, nil
}

func convertBackingStatus(dbStatus db.BackingStatus) model.BackingStatus {
	var status model.BackingStatus
	switch dbStatus {
	case db.BackingStatusPending:
		status = model.BackingStatusPending
	case db.BackingStatusRefunded:
		status = model.BackingStatusRefunded
	case db.BackingStatusReleased:
		status = model.BackingStatusReleased
	}
	return status
}
