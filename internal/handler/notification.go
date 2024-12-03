package handler

import (
	"context"
	ejson "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kanowfy/donorbox/internal/model"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Notification interface {
	GetNotificationsForUser(w http.ResponseWriter, r *http.Request)
	UpdateReadNotification(w http.ResponseWriter, r *http.Request)
	NotificationStreamHandler(w http.ResponseWriter, r *http.Request)
}

type notification struct {
	service   service.Notification
	notifChan chan model.Notification
	serverCtx context.Context
}

func NewNotification(service service.Notification, notifChan chan model.Notification, serverCtx context.Context) Notification {
	return &notification{
		service,
		notifChan,
		serverCtx,
	}
}

func (n *notification) GetNotificationsForUser(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	notifications, err := n.service.GetNotificationsForUser(r.Context(), uid)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"notifications": notifications,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (n *notification) UpdateReadNotification(w http.ResponseWriter, r *http.Request) {
	nid, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		httperror.NotFoundResponse(w, r)
		return
	}

	if err := n.service.UpdateReadNotification(r.Context(), nid); err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "notification updated",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (n *notification) NotificationStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	rc := http.NewResponseController(w)

	for {
		select {
		case <-n.serverCtx.Done():
			fmt.Printf("client disconnected")
			return
		case <-r.Context().Done():
			return
		case notif := <-n.notifChan:
			payload, _ := ejson.Marshal(notif)

			fmt.Fprintf(w, "data: %s\n\n", string(payload))

			if err := rc.Flush(); err != nil {
				log.Printf("error flushing events: %v", err)
				return
			}
		}

	}
}
