package publish

import "github.com/kanowfy/donorbox/internal/model"

type Publisher interface {
	Publish(notification model.Notification)
}

type publisher struct {
	notifChan chan model.Notification
}

func New(notifChan chan model.Notification) Publisher {
	return &publisher{
		notifChan,
	}
}

// Publish sends a notification to the configured notification channel
func (p *publisher) Publish(notification model.Notification) {
	p.notifChan <- notification
}
