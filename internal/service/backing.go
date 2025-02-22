package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
)

type Backing interface {
	CreateBacking(ctx context.Context, projectID int64, req dto.BackingRequest) error
	GetBackingsForProject(ctx context.Context, projectID int64) ([]model.Backing, error)
	GetProjectBackingStats(ctx context.Context, projectID int64) (*model.Backing, *model.Backing, *model.Backing, int64, error)
	GetBackingsForUser(ctx context.Context, userID int64) ([]dto.UserBackingResponse, error)
}

type backing struct {
	repository db.Querier
	auditSvc   AuditTrail
}

func NewBacking(repository db.Querier, auditSvc AuditTrail) Backing {
	return &backing{
		repository,
		auditSvc,
	}
}

func (b *backing) CreateBacking(ctx context.Context, projectID int64, req dto.BackingRequest) error {
	queries := b.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Create db backing
	backingParams := db.CreateBackingParams{
		ProjectID: projectID,
		Amount:    req.Amount,
	}

	if req.WordOfSupport != nil && *req.WordOfSupport != "" {
		backingParams.WordOfSupport = req.WordOfSupport
	}

	if req.UserID != nil {
		backingParams.UserID = req.UserID
	}

	backing, err := q.CreateBacking(ctx, backingParams)
	if err != nil {
		return err
	}

	if err := b.auditSvc.LogAction(ctx, LogActionParams{
		UserID:        backingParams.UserID,
		EntityType:    "backing",
		EntityID:      &backing.ID,
		OperationType: "CREATE",
		NewValue:      backing,
	}); err != nil {
		return err
	}

	// assuming the milestones are sorted
	milestones, err := q.GetMilestoneForProject(ctx, projectID)
	if err != nil {
		return err
	}

	remainingAmount := req.Amount
	for _, m := range milestones {
		backAmount := remainingAmount
		if m.CurrentFund < m.FundGoal {
			gap := m.FundGoal - m.CurrentFund
			if gap < backAmount {
				backAmount = gap
			}
			remainingAmount -= backAmount
			if err := q.UpdateMilestoneFund(ctx, db.UpdateMilestoneFundParams{
				ID:     m.ID,
				Amount: backAmount,
			}); err != nil {
				return err
			}

			if remainingAmount == 0 {
				break
			}
		}
	}

	// any remaining fund will go to the last milestone
	if remainingAmount > 0 {
		if err := q.UpdateMilestoneFund(ctx, db.UpdateMilestoneFundParams{
			ID:     milestones[len(milestones)-1].ID,
			Amount: remainingAmount,
		}); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

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
			CreatedAt:     convert.MustPgTimestampToTime(b.CreatedAt),
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

func (b *backing) GetProjectBackingStats(ctx context.Context, projectID int64) (*model.Backing, *model.Backing, *model.Backing, int64, error) {
	mostBacking, err := b.repository.GetMostBackingDonor(ctx, projectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, nil, 0, nil
		}
		return nil, nil, nil, 0, err
	}

	firstBacking, err := b.repository.GetFirstBackingDonor(ctx, projectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, nil, 0, nil
		}
		return nil, nil, nil, 0, err
	}

	recentBacking, err := b.repository.GetMostRecentBackingDonor(ctx, projectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, nil, 0, nil
		}
		return nil, nil, nil, 0, err
	}

	count, err := b.repository.GetBackingCountForProject(ctx, projectID)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	return &model.Backing{
			ID:            mostBacking.ID,
			ProjectID:     mostBacking.ProjectID,
			Amount:        mostBacking.Amount,
			WordOfSupport: mostBacking.WordOfSupport,
			CreatedAt:     convert.MustPgTimestampToTime(mostBacking.CreatedAt),
			Backer: &model.Backer{
				FirstName:      mostBacking.FirstName,
				LastName:       mostBacking.LastName,
				ProfilePicture: mostBacking.ProfilePicture,
			},
		}, &model.Backing{
			ID:            firstBacking.ID,
			ProjectID:     firstBacking.ProjectID,
			Amount:        firstBacking.Amount,
			WordOfSupport: firstBacking.WordOfSupport,
			CreatedAt:     convert.MustPgTimestampToTime(firstBacking.CreatedAt),
			Backer: &model.Backer{
				FirstName:      firstBacking.FirstName,
				LastName:       firstBacking.LastName,
				ProfilePicture: firstBacking.ProfilePicture,
			},
		}, &model.Backing{
			ID:            recentBacking.ID,
			ProjectID:     recentBacking.ProjectID,
			Amount:        recentBacking.Amount,
			WordOfSupport: recentBacking.WordOfSupport,
			CreatedAt:     convert.MustPgTimestampToTime(recentBacking.CreatedAt),
			Backer: &model.Backer{
				FirstName:      recentBacking.FirstName,
				LastName:       recentBacking.LastName,
				ProfilePicture: recentBacking.ProfilePicture,
			},
		}, count, nil
}

func (b *backing) GetBackingsForUser(ctx context.Context, userID int64) ([]dto.UserBackingResponse, error) {
	dbBackings, err := b.repository.GetBackingsForUser(ctx, &userID)
	if err != nil {
		return nil, fmt.Errorf("get backings for user: %w", err)
	}

	backings := make([]dto.UserBackingResponse, len(dbBackings))
	for i, b := range dbBackings {
		backings[i] = dto.UserBackingResponse{
			BackingID:           b.ID,
			ProjectID:           b.ProjectID,
			Amount:              b.Amount,
			WordOfSupport:       b.WordOfSupport,
			CreatedAt:           convert.MustPgTimestampToTime(b.CreatedAt),
			ProjectTitle:        b.Title,
			ProjectCoverPicture: b.CoverPicture,
		}
	}

	return backings, nil
}
