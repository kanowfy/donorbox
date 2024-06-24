package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/model"
)

type Transaction interface {
	GetAllTransactions(ctx context.Context) ([]model.Transaction, error)
	GetOneTransaction(ctx context.Context, txID uuid.UUID) (*model.Transaction, error)
	GetTransactionAudit(ctx context.Context, projectID uuid.UUID) ([]model.Transaction, error)
}

type transaction struct {
	repository db.Querier
}

func NewTransaction(repository db.Querier) Transaction {
	return &transaction{
		repository,
	}
}

func (t *transaction) GetAllTransactions(ctx context.Context) ([]model.Transaction, error) {
	dbTransactions, err := t.repository.GetAllTransactions(ctx)
	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction

	for _, t := range dbTransactions {
		transactions = append(transactions, model.Transaction{
			ID:              t.ID,
			ProjectID:       t.ProjectID,
			TransactionType: convertTransactionType(t.TransactionType),
			Amount:          0,
			InitiatorCardID: t.InitiatorCardID,
			RecipientCardID: t.RecipientCardID,
			Status:          convertTransactionStatus(t.Status),
			CreatedAt:       t.CreatedAt,
		})
	}

	return transactions, nil
}

func (t *transaction) GetOneTransaction(ctx context.Context, txID uuid.UUID) (*model.Transaction, error) {
	transaction, err := t.repository.GetTransactionByID(ctx, txID)
	if err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:              transaction.ID,
		ProjectID:       transaction.ProjectID,
		TransactionType: convertTransactionType(transaction.TransactionType),
		Amount:          transaction.Amount,
		InitiatorCardID: transaction.InitiatorCardID,
		RecipientCardID: transaction.RecipientCardID,
		Status:          convertTransactionStatus(transaction.Status),
		CreatedAt:       transaction.CreatedAt,
	}, nil
}

func (t *transaction) GetTransactionAudit(ctx context.Context, projectID uuid.UUID) ([]model.Transaction, error) {
	dbTransactions, err := t.repository.GetTransactionsForProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction

	for _, t := range dbTransactions {
		transactions = append(transactions, model.Transaction{
			ID:              t.ID,
			ProjectID:       t.ProjectID,
			TransactionType: convertTransactionType(t.TransactionType),
			Amount:          0,
			InitiatorCardID: t.InitiatorCardID,
			RecipientCardID: t.RecipientCardID,
			Status:          convertTransactionStatus(t.Status),
			CreatedAt:       t.CreatedAt,
		})
	}

	return transactions, nil
}

func convertTransactionStatus(dbStatus db.TransactionStatus) model.TransactionStatus {
	var status model.TransactionStatus
	switch dbStatus {
	case db.TransactionStatusPending:
		status = model.TransactionStatusPending
	case db.TransactionStatusCompleted:
		status = model.TransactionStatusCompleted
	case db.TransactionStatusFailed:
		status = model.TransactionStatusFailed
	}
	return status
}

func convertTransactionType(dbType db.TransactionType) model.TransactionType {
	var typ model.TransactionType
	switch dbType {
	case db.TransactionTypeBacking:
		typ = model.TransactionTypeBacking
	case db.TransactionTypePayout:
		typ = model.TransactionTypePayout
	case db.TransactionTypeRefund:
		typ = model.TransactionTypeRefund
	}
	return typ
}
