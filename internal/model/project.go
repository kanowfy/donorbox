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
	ProjectStatusDisputed ProjectStatus = "disputed"
	ProjectStatusStopped  ProjectStatus = "stopped"
)

type MilestoneStatus string

const (
	MilestoneStatusPending      MilestoneStatus = "pending"
	MilestoneStatusFundReleased MilestoneStatus = "fund_released"
	MilestoneStatusCompleted    MilestoneStatus = "completed"
	MilestoneStatusRefuted      MilestoneStatus = "refuted"
)

type ProofStatus string

const (
	ProofStatusPending  ProofStatus = "pending"
	ProofStatusApproved ProofStatus = "approved"
	ProofStatusRejected ProofStatus = "rejected"
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
	Status          MilestoneStatus      `json:"status"`
	Completion      *MilestoneCompletion `json:"milestone_completion,omitempty"`
	SpendingProofs  []SpendingProof      `json:"spending_proofs,omitempty"`
}

type MilestoneCompletion struct {
	TransferAmount int64     `json:"transfer_amount"`
	TransferNote   *string   `json:"transfer_note,omitempty"`
	TransferImage  *string   `json:"transfer_image,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type SpendingProof struct {
	ID            int64       `json:"id"`
	TransferImage string      `json:"transfer_image"`
	Description   string      `json:"description"`
	ProofMedia    string      `json:"proof_media"`
	Status        ProofStatus `json:"status"`
	RejectedCause *string     `json:"rejected_cause,omitempty"`
	CreatedAt     time.Time   `json:"created_at"`
}

type ReportStatus string

const (
	ReportStatusPending   ReportStatus = "pending"
	ReportStatusDismissed ReportStatus = "dismissed"
	ReportStatusResolved  ReportStatus = "resolved"
)

type ProjecReport struct {
	ID          int64        `json:"id"`
	ProjectID   int64        `json:"project_id"`
	Email       string       `json:"email"`
	FullName    string       `json:"full_name"`
	PhoneNumber string       `json:"phone_number"`
	Relation    *string      `json:"relation,omitempty"`
	Reason      string       `json:"reason"`
	Details     string       `json:"details"`
	ProofImage  *string      `json:"proof_image,omitempty"`
	Status      ReportStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
}
