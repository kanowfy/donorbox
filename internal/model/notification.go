package model

import "time"

type NotificationType string

type Notification struct {
	ID               int64            `json:"id"`
	UserID           int64            `json:"user_id"`
	NotificationType NotificationType `json:"type"`
	Message          string           `json:"message"`
	ProjectID        *int64           `json:"project_id,omitempty"`
	IsRead           bool             `json:"is_read"`
	CreatedAt        time.Time        `json:"created_at"`
}
