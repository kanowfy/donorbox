package handler

import (
	"net/http"

	"github.com/kanowfy/donorbox/pkg/httperror"
	"github.com/kanowfy/donorbox/pkg/json"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status":  "up",
		"version": "0.0.1",
	}

	err := json.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		httperror.ServerErrorResponse(w, r, err)
	}
}
