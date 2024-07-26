package model

import (
	"time"

	"github.com/google/uuid"
)

type ProjectStatus string

const (
	ProjectStatusOngoing         ProjectStatus = "ongoing"
	ProjectStatusEnded           ProjectStatus = "ended"
	ProjectStatusCompletedPayout ProjectStatus = "completed_payout"
	ProjectStatusCompletedRefund ProjectStatus = "completed_refund"
)

type Project struct {
	ID            uuid.UUID     `json:"id"`
	UserID        uuid.UUID     `json:"user_id"`
	CategoryID    int32         `json:"category_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	CoverPicture  string        `json:"cover_picture"`
	GoalAmount    int64         `json:"goal_amount"`
	CurrentAmount int64         `json:"current_amount"`
	Country       string        `json:"country"`
	Province      string        `json:"province"`
	CardID        uuid.UUID     `json:"card_id"`
	StartDate     time.Time     `json:"start_date"`
	EndDate       time.Time     `json:"end_date"`
	Status        ProjectStatus `json:"status"`
	BackingCount  *int64        `json:"backing_count,omitempty"`
}

type ProjectUpdate struct {
	ID              uuid.UUID
	ProjectID       uuid.UUID
	AttachmentPhoto *string
	Description     string
	CreatedAt       time.Time
}
