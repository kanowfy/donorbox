package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthz", nil)
	recorder := httptest.NewRecorder()

	Healthcheck(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	want := map[string]string{
		"status":  "up",
		"version": "0.0.1",
	}

	var m map[string]string
	err = json.Unmarshal(body, &m)
	assert.NoError(t, err)

	if !reflect.DeepEqual(m, want) {
		t.Errorf("Unexpected body returned, want %v got %v", want, m)
	}
}
