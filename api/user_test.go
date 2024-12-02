package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"profile-creation/types"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	app := fiber.New()
	app.Post("/user", CreateUser)

	tests := []struct {
		name           string
		payload        types.UserRequest
		expectedStatus int
	}{
		{
			name: "Valid User Input",
			payload: types.UserRequest{
				Pan:    "ABCDE1234F",
				Mobile: "9876543210",
				Email:  "test@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid PAN Format",
			payload: types.UserRequest{
				Pan:    "ABCDE123",
				Mobile: "9876543210",
				Email:  "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Mobile Length",
			payload: types.UserRequest{
				Pan:    "ABCDE1234F",
				Mobile: "12345",
				Email:  "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Email Format",
			payload: types.UserRequest{
				Pan:    "ABCDE1234F",
				Mobile: "9876543210",
				Email:  "invalid-email",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)

			req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
