package handler

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type AuditTrail interface {
	GetAuditHistory(w http.ResponseWriter, r *http.Request)
}

type auditTrail struct {
	service service.AuditTrail
}

func NewAuditTrail(service service.AuditTrail) AuditTrail {
	return &auditTrail{
		service,
	}
}

func (a *auditTrail) GetAuditHistory(w http.ResponseWriter, r *http.Request) {
	trails, err := a.service.GetAuditHistory(r.Context())
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"audits": trails,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
