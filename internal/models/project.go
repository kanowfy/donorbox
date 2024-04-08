package models

import "time"

type CreateProjectRequest struct {
	UserID       string    `json:"user_id" validate:"required,uuid4"`
	Title        string    `json:"title" validate:"required,min=8"`
	Description  string    `json:"description" validate:"required,min=50"`
	CoverPicture string    `json:"cover_picture" validate:"required,http_url"`
	GoalAmount   string    `json:"goal_amount" validate:"required,numeric"`
	Country      string    `json:"country" validate:"required"`
	Province     string    `json:"province" validate:"required"`
	EndDate      time.Time `json:"end_date" validate:"required"`
}

type UpdateProjectRequest struct {
	Title        *string    `json:"title,omitempty" validate:"omitnil,min=8"`
	Description  *string    `json:"description,omitempty" validate:"omitnil,min=50"`
	CoverPicture *string    `json:"cover_picture,omitempty" validate:"omitnil,http_url"`
	GoalAmount   *string    `json:"goal_amount,omitempty" validate:"omitnil,numeric"`
	Country      *string    `json:"country,omitempty"`
	Province     *string    `json:"province,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
}
