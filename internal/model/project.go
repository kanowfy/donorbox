package model

import (
	"time"
)

type ProjectStatus string

const (
	ProjectStatusPending  ProjectStatus = "pending"
	ProjectStatusOngoing  ProjectStatus = "ongoing"
	ProjectStatusRejected ProjectStatus = "rejected"
	ProjectStatusFinished ProjectStatus = "finished"
)

type Project struct {
	ID             int64         `json:"id"`
	UserID         int64         `json:"user_id"`
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
	ID              int64                `json:"id"`
	ProjectID       int64                `json:"project_id"`
	Title           string               `json:"title"`
	Description     *string              `json:"description,omitempty"`
	FundGoal        int64                `json:"fund_goal"`
	CurrentFund     int64                `json:"current_fund"`
	BankDescription string               `json:"bank_description"`
	Completed       bool                 `json:"completed"`
	Completion      *MilestoneCompletion `json:"milestone_completion,omitempty"`
}

type MilestoneCompletion struct {
	TransferAmount int64     `json:"transfer_amount"`
	TransferNote   *string   `json:"transfer_note,omitempty"`
	TransferImage  *string   `json:"transfer_image,omitempty"`
	CompletedAt    time.Time `json:"completed_at"`
}

type ProjectUpdate struct {
	ID              int64
	ProjectID       int64
	AttachmentPhoto *string
	Description     string
	CreatedAt       time.Time
}
