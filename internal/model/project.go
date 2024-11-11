package model

import (
	"time"

	"github.com/google/uuid"
)

type ProjectStatus string

const (
	ProjectStatusPending  ProjectStatus = "pending"
	ProjectStatusOngoing  ProjectStatus = "ongoing"
	ProjectStatusRejected ProjectStatus = "rejected"
	ProjectStatusFinished ProjectStatus = "finished"
)

type Project struct {
	ID             uuid.UUID     `json:"id"`
	UserID         uuid.UUID     `json:"user_id"`
	CategoryID     int32         `json:"category_id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	CoverPicture   string        `json:"cover_picture"`
	ReceiverName   string        `json:"receiver_name"`
	ReceiverNumber string        `json:"receiver_number"`
	Address        string        `json:"address"`
	District       string        `json:"district"`
	City           string        `json:"city"`
	Country        string        `json:"country"`
	StartDate      time.Time     `json:"start_date"`
	EndDate        time.Time     `json:"end_date"`
	Status         ProjectStatus `json:"status"`
	BackingCount   *int64        `json:"backing_count,omitempty"`
}

type ProjectUpdate struct {
	ID              uuid.UUID
	ProjectID       uuid.UUID
	AttachmentPhoto *string
	Description     string
	CreatedAt       time.Time
}
