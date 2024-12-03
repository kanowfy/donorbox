package dto

import (
	"time"

	"github.com/kanowfy/donorbox/internal/model"
)

type ProjectApprovalRequest struct {
	ProjectID    int64   `json:"project_id" validate:"required,number"`
	Approved     *bool   `json:"approved,omitempty" validate:"omitnil"`
	RejectReason *string `json:"reject_reason,omitempty" validate:"omitnil"`
}

type UnresolvedMilestoneDto struct {
	Milestone      model.Milestone `json:"milestone"`
	Address        string          `json:"address"`
	District       string          `json:"district"`
	City           string          `json:"city"`
	Country        string          `json:"country"`
	ReceiverName   string          `json:"receiver_name"`
	ReceiverNumber string          `json:"receiver_number"`
}

type ResolveMilestoneRequest struct {
	Amount      int64   `json:"amount" validate:"required,number"`
	Description *string `json:"description" validate:"required"`
	Image       *string `json:"image,omitempty" validate:"omitnil"`
}

type PendingUserVerificationResponse struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DocumentUrl string    `json:"document_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type VerificationApprovalRequest struct {
	UserID       int64   `json:"user_id" validate:"required,number"`
	Approved     *bool   `json:"approved,omitempty" validate:"omitnil"`
	RejectReason *string `json:"reject_reason,omitempty" validate:"omitnil"`
}

type GetStatisticsResponse struct {
	TotalFund                int64           `json:"total_fund"`
	DonationCount            int64           `json:"donation_count"`
	ProjectCount             ProjectCount    `json:"project_count"`
	PendingVerificationCount int64           `json:"verification_count"`
	CategoriesCount          []CategoryCount `json:"categories_count"`
}

type ProjectCount struct {
	Pending  int64 `json:"pending"`
	Ongoing  int64 `json:"ongoing"`
	Finished int64 `json:"finished"`
	Rejected int64 `json:"rejected"`
}

type CategoryCount struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
