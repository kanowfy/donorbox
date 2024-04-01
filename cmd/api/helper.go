package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data map[string]interface{}, headers http.Header) error {
	encoded, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	encoded = append(encoded, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(encoded)

	return nil
}
