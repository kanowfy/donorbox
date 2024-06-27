package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kanowfy/donorbox/internal/config"
	"github.com/kanowfy/donorbox/internal/dto"
	"github.com/kanowfy/donorbox/internal/service"
	"github.com/kanowfy/donorbox/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type mockOutput struct {
		token string
		err   error
	}

	cases := []struct {
		name string
		req  dto.LoginRequest
		// whether we expect the handler would call service method, to prevent testify asking us to call mock method
		useMock        bool
		mockOut        mockOutput
		expectedStatus int
		//expectedBody map[string]interface{}
	}{
		{
			"success",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678",
			},
			true,
			mockOutput{
				token: "12345678abc",
				err:   nil,
			},
			http.StatusOK,
		},
		{
			"failed validation",
			dto.LoginRequest{
				Email:    "abcgmail.com",
				Password: "12345678",
			},
			false,
			// here we expect the code wouldn't even reach the service call, so anything here would work
			mockOutput{
				token: "",
				err:   nil,
			},
			http.StatusUnprocessableEntity,
		},
		{
			"user not found",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678",
			},
			true,
			mockOutput{
				token: "",
				err:   service.ErrUserNotFound,
			},
			http.StatusNotFound,
		},
		{
			"wrong password",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678",
			},
			true,
			mockOutput{
				token: "",
				err:   service.ErrWrongPassword,
			},
			http.StatusUnauthorized,
		},
		{
			"server error",
			dto.LoginRequest{
				Email:    "abc@gmail.com",
				Password: "12345678",
			},
			true,
			mockOutput{
				token: "",
				err:   errors.New("random error"),
			},
			http.StatusInternalServerError,
		},
	}

	// validator is threadsafe so we can reuse this single instance accross tests
	validator := validator.New()

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			authMock := mocks.NewAuth(t)
			if tt.useMock {
				authMock.EXPECT().Login(context.Background(), tt.req).Return(tt.mockOut.token, tt.mockOut.err)
			}

			authHandler := NewAuth(authMock, validator, config.Config{})

			recorder := httptest.NewRecorder()
			b, _ := json.Marshal(tt.req)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewReader(b))

			authHandler.Login(recorder, req)

			resp := recorder.Result()
			defer resp.Body.Close()

			_, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Unexpected status code returned, want '%d' got '%d'", tt.expectedStatus, resp.StatusCode)
			}
		})
	}

}
