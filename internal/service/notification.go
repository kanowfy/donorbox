package service

import (
	"context"

	"github.com/kanowfy/donorbox/internal/convert"
	"github.com/kanowfy/donorbox/internal/db"
	"github.com/kanowfy/donorbox/internal/model"
)

type Notification interface {
	GetNotificationsForUser(ctx context.Context, userID int64) ([]model.Notification, error)
	UpdateReadNotification(ctx context.Context, id int64) error
}

type notification struct {
	repository db.Querier
}

func NewNotification(querier db.Querier) Notification {
	return &notification{
		repository: querier,
	}
}

func (n *notification) GetNotificationsForUser(ctx context.Context, userID int64) ([]model.Notification, error) {
	dbNotifications, err := n.repository.GetNotificationsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var notifications []model.Notification

	for _, notifs := range dbNotifications {
		notifications = append(notifications, model.Notification{
			ID:               notifs.ID,
			UserID:           notifs.UserID,
			NotificationType: model.NotificationType(notifs.NotificationType),
			Message:          notifs.Message,
			ProjectID:        notifs.ProjectID,
			IsRead:           notifs.IsRead,
			CreatedAt:        convert.MustPgTimestampToTime(notifs.CreatedAt),
		})
	}

	return notifications, nil
}

func (n *notification) UpdateReadNotification(ctx context.Context, id int64) error {
	if err := n.repository.UpdateReadNotification(ctx, id); err != nil {
		return err
	}

	return nil
}
