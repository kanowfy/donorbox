package models

import "time"

type CreateProjectRequest struct {
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CoverPicture string    `json:"cover_picture"`
	GoalAmount   string    `json:"goal_amount"`
	Country      string    `json:"country"`
	Province     string    `json:"province"`
	EndDate      time.Time `json:"end_date"`
}
