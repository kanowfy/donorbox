package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/model"
)

type AuditTrail interface {
	LogAction(ctx context.Context, params LogActionParams) error
	GetAuditHistory(ctx context.Context) ([]model.AuditTrail, error)
}

type auditTrail struct {
	repository db.Querier
}

func NewAuditTrail(repository db.Querier) AuditTrail {
	return &auditTrail{
		repository,
	}
}

type LogActionParams struct {
	UserID        *int64
	EscrowID      *int64
	EntityType    string
	EntityID      *int64
	OperationType string
	FieldName     string
	OldValue      any
	NewValue      any
}

func (a *auditTrail) LogAction(ctx context.Context, params LogActionParams) error {
	dbParams := db.CreateAuditLogParams{
		UserID:        params.UserID,
		EscrowID:      params.EscrowID,
		EntityType:    params.EntityType,
		EntityID:      params.EntityID,
		OperationType: params.OperationType,
		FieldName:     params.FieldName,
	}

	if params.OldValue != nil {
		jsonVal, err := json.Marshal(params.OldValue)
		if err != nil {
			return fmt.Errorf("log action: %w", err)
		}

		dbParams.OldValue = jsonVal
	}

	if params.NewValue != nil {
		jsonVal, err := json.Marshal(params.NewValue)
		if err != nil {
			return fmt.Errorf("log action: %w", err)
		}

		dbParams.NewValue = jsonVal
	}

	// leaves off return value for now
	if _, err := a.repository.CreateAuditLog(ctx, dbParams); err != nil {
		return fmt.Errorf("log action: %w", err)
	}

	return nil
}

func (a *auditTrail) GetAuditHistory(ctx context.Context) ([]model.AuditTrail, error) {
	dbTrails, err := a.repository.GetAuditHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("get audit history: %w", err)
	}

	var trails []model.AuditTrail
	for _, t := range dbTrails {
		trail := model.AuditTrail{
			ID:            t.ID,
			UserID:        t.UserID,
			EscrowID:      t.EscrowID,
			EntityType:    t.EntityType,
			EntityID:      t.EntityID,
			OperationType: t.OperationType,
			FieldName:     t.FieldName,
			CreatedAt:     convert.MustPgTimestampToTime(t.CreatedAt),
		}

		if err = json.Unmarshal(t.OldValue, &trail.OldValue); err != nil {
			if errors.Is(err, &json.InvalidUnmarshalError{}) {
				trail.OldValue = nil
			}
		}

		if err = json.Unmarshal(t.NewValue, &trail.NewValue); err != nil {
			if errors.Is(err, &json.InvalidUnmarshalError{}) {
				trail.NewValue = nil
			}
		}

		trails = append(trails, trail)
	}

	return trails, nil
}
