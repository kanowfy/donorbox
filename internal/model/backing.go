package model

import (
	"time"

	"github.com/google/uuid"
)

type BackingStatus string

const (
	BackingStatusPending  BackingStatus = "pending"
	BackingStatusReleased BackingStatus = "released"
	BackingStatusRefunded BackingStatus = "refunded"
)

type Backing struct {
	ID              uuid.UUID     `json:"id"`
	ProjectID       uuid.UUID     `json:"project_id"`
	Amount          int64         `json:"amount"`
	WordOfSupport   *string       `json:"word_of_support,omitempty"`
	Status          BackingStatus `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
	BackerFirstName *string       `json:"first_name,omitempty"`
	BackerLastName  *string       `json:"last_name,omitempty"`
	ProfilePicture  *string       `json:"profile_picture,omitempty"`
}
