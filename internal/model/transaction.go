package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	TransactionType   string
	TransactionStatus string
)

const (
	TransactionTypeBacking TransactionType = "backing"
	TransactionTypePayout  TransactionType = "payout"
	TransactionTypeRefund  TransactionType = "refund"
)

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID              uuid.UUID         `json:"id"`
	ProjectID       uuid.UUID         `json:"project_id"`
	TransactionType TransactionType   `json:"transaction_type"`
	Amount          int64             `json:"amount"`
	InitiatorCardID uuid.UUID         `json:"initiator_card_id"`
	RecipientCardID uuid.UUID         `json:"recipient_card_id"`
	Status          TransactionStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
}
