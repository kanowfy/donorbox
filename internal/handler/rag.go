package handler

import (
	"net/http"

	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

type Rag interface {
	AddDocuments(w http.ResponseWriter, r *http.Request)
	Query(w http.ResponseWriter, r *http.Request)
}

type rag struct {
	service service.Rag
}

func NewRag(service service.Rag) Rag {
	return &rag{
		service,
	}
}

func (rg *rag) AddDocuments(w http.ResponseWriter, r *http.Request) {
	var req dto.AddDocumentRequest
	if err := json.ReadJSON(w, r, &req); err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	if err := rg.service.AddDocuments(r.Context(), req); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}

	if err := json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "documents added",
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}

func (rg *rag) Query(w http.ResponseWriter, r *http.Request) {
	var req dto.QueryRequest
	if err := json.ReadJSON(w, r, &req); err != nil {
		httperror.BadRequestResponse(w, r, err)
		return
	}

	res, err := rg.service.Query(r.Context(), req)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
		return
	}

	if err = json.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"answer": res,
	}, nil); err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
