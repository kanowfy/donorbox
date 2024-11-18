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
	FundGoal       int64         `json:"fund_goal"`
	TotalFund      int64         `json:"total_fund"`
	CoverPicture   string        `json:"cover_picture"`
	ReceiverName   string        `json:"receiver_name"`
	ReceiverNumber string        `json:"receiver_number"`
	Address        string        `json:"address"`
	District       string        `json:"district"`
	City           string        `json:"city"`
	Country        string        `json:"country"`
	EndDate        time.Time     `json:"end_date"`
	Status         ProjectStatus `json:"status"`
	BackingCount   *int64        `json:"backing_count,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
}

type Milestone struct {
	ID              uuid.UUID `json:"id"`
	ProjectID       uuid.UUID `json:"project_id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description,omitempty"`
	FundGoal        int64     `json:"fund_goal"`
	CurrentFund     int64     `json:"current_fund"`
	BankDescription string    `json:"bank_description"`
	Completed       bool      `json:"completed"`
}

type ProjectUpdate struct {
	ID              uuid.UUID
	ProjectID       uuid.UUID
	AttachmentPhoto *string
	Description     string
	CreatedAt       time.Time
}
