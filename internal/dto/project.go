package dto

import (
	"time"

	"github.com/kanowfy/donorbox/internal/model"
)

type CreateProjectRequest struct {
	CategoryID     int                `json:"category_id" validate:"required,min=1,max=7"`
	Title          string             `json:"title" validate:"required,min=8"`
	Description    string             `json:"description" validate:"required,min=50"`
	Milestones     []MilestoneRequest `json:"milestones" validate:"required"`
	CoverPicture   string             `json:"cover_picture" validate:"required,http_url"`
	EndDate        time.Time          `json:"end_date" validate:"required"`
	ReceiverNumber string             `json:"receiver_number" validate:"required"`
	ReceiverName   string             `json:"receiver_name" validate:"required"`
	Address        string             `json:"address" validate:"required"`
	District       string             `json:"district" validate:"required"`
	City           string             `json:"city" validate:"required"`
	Country        string             `json:"country" validate:"required"`
}

type MilestoneRequest struct {
	Title           string  `json:"title" validate:"required,min=8"`
	Description     *string `json:"description,omitempty" validate:"omitnil,min=10"`
	FundGoal        string  `json:"fund_goal" validate:"required,number"`
	BankDescription string  `json:"bank_description" validate:"required"`
}

type CreateProjectResponse struct {
	Project    model.Project     `json:"project" validate:"required"`
	Milestones []model.Milestone `json:"milestones" validate:"required"`
}

type UpdateProjectRequest struct {
	Title          *string    `json:"title,omitempty" validate:"omitnil,min=8"`
	Description    *string    `json:"description,omitempty" validate:"omitnil,min=50"`
	CoverPicture   *string    `json:"cover_picture,omitempty" validate:"omitnil,http_url"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	ReceiverNumber *string    `json:"receiver_number" validate:"required"`
	ReceiverName   *string    `json:"receiver_name" validate:"required"`
	Address        *string    `json:"address" validate:"required"`
	District       *string    `json:"district" validate:"required"`
	City           *string    `json:"city" validate:"required"`
	Country        *string    `json:"country" validate:"required"`
}

type CreateProjectUpdateRequest struct {
	ProjectID       string  `json:"project_id" validate:"required,number"`
	AttachmentPhoto *string `json:"attachment_photo" validate:"omitnil,http_url"`
	Description     string  `json:"description" validate:"required"`
}

type CreateProjectCommentRequest struct {
	ProjectID       string  `json:"project_id" validate:"required,number"`
	BackerID        string  `json:"backer_id" validate:"required,number"`
	ParentCommentID *string `json:"parent_comment_id,omitempty" validate:"omitnil,number"`
	Content         string  `json:"content" validate:"required"`
}

type ProjectApprovalRequest struct {
	ProjectID    string  `json:"project_id" validate:"required,number"`
	Approved     bool    `json:"approved" validate:"required"`
	RejectReason *string `json:"reject_reason,omitempty" validate:"omitnil"`
	Description  *string `json:"description,omitempty" validate:"omitnil"`
}
