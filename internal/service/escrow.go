package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type Escrow interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
	GetEscrowByID(ctx context.Context, id uuid.UUID) (*model.EscrowUser, error)
	Payout(ctx context.Context, projectID uuid.UUID, escrow *model.EscrowUser) error
	Refund(ctx context.Context, projectID uuid.UUID, escrow *model.EscrowUser) error
	GetStatistics(ctx context.Context) (*ProjectStatistics, []WeeklyTransactions, error)
}

type escrow struct {
	repository db.Querier
}

func NewEscrow(querier db.Querier) Escrow {
	return &escrow{
		repository: querier,
	}
}

func (e *escrow) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	escrow, err := e.repository.GetEscrowUserByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrUserNotFound
	}

	// validate password
	if err = bcrypt.CompareHashAndPassword([]byte(escrow.HashedPassword), []byte(req.Password)); err != nil {
		return "", ErrWrongPassword
	}

	token, err := token.GenerateToken(escrow.ID.String(), time.Hour*3*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (e *escrow) GetEscrowByID(ctx context.Context, id uuid.UUID) (*model.EscrowUser, error) {
	escrow, err := e.repository.GetEscrowUserByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return &model.EscrowUser{
		ID:        escrow.ID,
		Email:     escrow.Email,
		UserType:  model.ESCROW,
		CardID:    escrow.CardID,
		CreatedAt: escrow.CreatedAt,
	}, nil
}

func (e *escrow) Payout(ctx context.Context, projectID uuid.UUID, escrow *model.EscrowUser) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	/* flow for payout
	projectID -> set all backing releated to project to released, set project status to completed_payout
	-> create a new transaction
	*/

	if err = q.UpdateProjectBackingStatus(ctx, db.UpdateProjectBackingStatusParams{
		ProjectID: projectID,
		Status:    db.BackingStatusReleased,
	}); err != nil {
		return err
	}

	if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
		ID:     projectID,
		Status: db.ProjectStatusCompletedPayout,
	}); err != nil {
		return err
	}

	project, err := q.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}

	_, err = q.CreateTransaction(ctx, db.CreateTransactionParams{
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

func (e *escrow) Refund(ctx context.Context, projectID uuid.UUID, escrow *model.EscrowUser) error {
	queries := e.repository.(*db.Queries)
	q, tx, err := queries.BeginTX(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	/* flow for refund
	projectID -> set all backing releated to project to refuned, set project status to completed_refund
	-> create a new transaction
	*/

	if err = q.UpdateProjectBackingStatus(ctx, db.UpdateProjectBackingStatusParams{
		ProjectID: projectID,
		Status:    db.BackingStatusRefunded,
	}); err != nil {
		return err
	}

	if err = q.UpdateProjectStatus(ctx, db.UpdateProjectStatusParams{
		ID:     projectID,
		Status: db.ProjectStatusCompletedRefund,
	}); err != nil {
		return err
	}

	transactions, err := q.GetBackingTransactionsForProject(ctx, projectID)
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		_, err = q.CreateTransaction(ctx, db.CreateTransactionParams{
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

type ProjectStatistics struct {
	Ended           int
	Ongoing         int
	CompletedPayout int
	CompletedRefund int
	Balance         int // should not be here?
}

type WeeklyTransactions struct {
	Week     time.Time
	Backings int
	Payouts  int
	Refunds  int
}

func (e *escrow) GetStatistics(ctx context.Context) (*ProjectStatistics, []WeeklyTransactions, error) {
	dbStats, err := e.repository.GetStatistics(ctx)
	if err != nil {
		return nil, nil, err
	}

	dbTransactions, err := e.repository.GetTransactionsStatsByWeek(ctx)
	if err != nil {
		return nil, nil, err
	}

	stats := &ProjectStatistics{
		Ended:           int(dbStats.Ended),
		Ongoing:         int(dbStats.Ongoing),
		CompletedPayout: int(dbStats.CompletedPayout),
		CompletedRefund: int(dbStats.CompletedRefund),
		Balance:         int(dbStats.Balance),
	}

	var weeklyTxs []WeeklyTransactions
	for _, tx := range dbTransactions {
		weeklyTxs = append(weeklyTxs, WeeklyTransactions{
			Week:     tx.Week.Time,
			Backings: int(tx.Backings),
			Payouts:  int(tx.Payouts),
			Refunds:  int(tx.Refunds),
		})
	}

	return stats, weeklyTxs, nil
}
