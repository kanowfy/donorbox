package model

import "time"

type AuditTrail struct {
	ID            int64     `json:"id"`
	UserID        *int64    `json:"user_id,omitempty"`
	EscrowID      *int64    `json:"escrow_id,omitempty"`
	EntityType    string    `json:"entity_type"`
	EntityID      *int64    `json:"entity_id,omitempty"`
	OperationType string    `json:"operation_type"`
	FieldName     string    `json:"field_name"`
	OldValue      any       `json:"old_value"`
	NewValue      any       `json:"new_value"`
	CreatedAt     time.Time `json:"created_at"`
}
