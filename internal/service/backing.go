package service

import (
	"context"
	"strconv"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
)

type Backing interface {
	AcceptBacking(ctx context.Context, projectID int64, req dto.BackingRequest) error
	GetBackingsForProject(ctx context.Context, projectID int64) ([]model.Backing, error)
	GetProjectBackingAggregation(ctx context.Context, projectID int64) (*BackingAggregation, error)
}

type backing struct {
	repository db.Querier
}

func NewBacking(repository db.Querier) Backing {
	return &backing{
		repository,
	}
}

func (b *backing) AcceptBacking(ctx context.Context, projectID int64, req dto.BackingRequest) error {
	amount := convert.MustStringToInt64(req.Amount)
	// Stripe, card...

	// Create db backing
	backingParams := db.CreateBackingParams{
		ProjectID:     projectID,
		Amount:        amount,
		WordOfSupport: req.WordOfSupport,
	}

	if req.UserID != nil {
		uid, err := strconv.ParseInt(*req.UserID, 10, 64)
		if err != nil {
			return err
		}

		backingParams.UserID = &uid
	}

	_, err := b.repository.CreateBacking(ctx, backingParams)
	if err != nil {
		return err
	}

	// Update project, milestone funds
	milestone, err := b.repository.GetCurrentMilestone(ctx, projectID)
	if err != nil {
		return err
	}

	if err := b.repository.UpdateMilestoneFund(ctx, db.UpdateMilestoneFundParams{
		ID:     milestone.ID,
		Amount: amount,
	}); err != nil {
		return err
	}

	if err := b.repository.UpdateProjectFund(ctx, db.UpdateProjectFundParams{
		ID:     projectID,
		Amount: amount,
	}); err != nil {
		return err
	}

	// Blockchain stuff

	return nil
}

/*
func (b *backing) AcceptBacking(ctx context.Context, projectID int64, request dto.BackingRequest) error {
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
*/

func (b *backing) GetBackingsForProject(ctx context.Context, projectID int64) ([]model.Backing, error) {
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
			ID:            b.ID,
			ProjectID:     b.ProjectID,
			Amount:        b.Amount,
			WordOfSupport: b.WordOfSupport,
			CreatedAt:     b.CreatedAt,
			Backer: &model.Backer{
				FirstName:      b.FirstName,
				LastName:       b.LastName,
				ProfilePicture: b.ProfilePicture,
			},
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

func (b *backing) GetProjectBackingAggregation(ctx context.Context, projectID int64) (*BackingAggregation, error) {
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
			ID:            mostBacking.ID,
			ProjectID:     mostBacking.ProjectID,
			Amount:        mostBacking.Amount,
			WordOfSupport: mostBacking.WordOfSupport,
			CreatedAt:     mostBacking.CreatedAt,
			Backer: &model.Backer{
				FirstName:      mostBacking.FirstName,
				LastName:       mostBacking.LastName,
				ProfilePicture: mostBacking.ProfilePicture,
			},
		},
		FirstBacking: model.Backing{
			ID:            firstBacking.ID,
			ProjectID:     firstBacking.ProjectID,
			Amount:        firstBacking.Amount,
			WordOfSupport: firstBacking.WordOfSupport,
			CreatedAt:     firstBacking.CreatedAt,
			Backer: &model.Backer{
				FirstName:      firstBacking.FirstName,
				LastName:       firstBacking.LastName,
				ProfilePicture: firstBacking.ProfilePicture,
			},
		},
		RecentBacking: model.Backing{
			ID:            recentBacking.ID,
			ProjectID:     recentBacking.ProjectID,
			Amount:        recentBacking.Amount,
			WordOfSupport: recentBacking.WordOfSupport,
			CreatedAt:     recentBacking.CreatedAt,
			Backer: &model.Backer{
				FirstName:      recentBacking.FirstName,
				LastName:       recentBacking.LastName,
				ProfilePicture: recentBacking.ProfilePicture,
			},
		},
		TotalBacking: count,
	}

	return backingAgg, nil
}
