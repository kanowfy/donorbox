package dto

import "time"

type CreateProjectRequest struct {
	CategoryID     int       `json:"category_id" validate:"required,min=1,max=7"`
	Title          string    `json:"title" validate:"required,min=8"`
	Description    string    `json:"description" validate:"required,min=50"`
	CoverPicture   string    `json:"cover_picture" validate:"required,http_url"`
	GoalAmount     string    `json:"goal_amount" validate:"required,number"`
	EndDate        time.Time `json:"end_date" validate:"required"`
	ReceiverNumber string    `json:"receiver_number" validate:"required"`
	ReceiverName   string    `json:"receiver_name" validate:"required"`
	Address        string    `json:"address" validate:"required"`
	District       string    `json:"district" validate:"required"`
	City           string    `json:"city" validate:"required"`
	Country        string    `json:"country" validate:"required"`
}

type UpdateProjectRequest struct {
	Title          *string    `json:"title,omitempty" validate:"omitnil,min=8"`
	Description    *string    `json:"description,omitempty" validate:"omitnil,min=50"`
	CoverPicture   *string    `json:"cover_picture,omitempty" validate:"omitnil,http_url"`
	GoalAmount     *string    `json:"goal_amount,omitempty" validate:"omitnil,number"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	ReceiverNumber *string    `json:"receiver_number" validate:"required"`
	ReceiverName   *string    `json:"receiver_name" validate:"required"`
	Address        *string    `json:"address" validate:"required"`
	District       *string    `json:"district" validate:"required"`
	City           *string    `json:"city" validate:"required"`
	Country        *string    `json:"country" validate:"required"`
}

type CreateProjectUpdateRequest struct {
	ProjectID       string  `json:"project_id" validate:"required,uuid4"`
	AttachmentPhoto *string `json:"attachment_photo" validate:"omitnil,http_url"`
	Description     string  `json:"description" validate:"required"`
}

type CreateProjectCommentRequest struct {
	ProjectID       string  `json:"project_id" validate:"required,uuid4"`
	BackerID        string  `json:"backer_id" validate:"required,uuid4"`
	ParentCommentID *string `json:"parent_comment_id,omitempty" validate:"omitnil,uuid4"`
	Content         string  `json:"content" validate:"required"`
}
