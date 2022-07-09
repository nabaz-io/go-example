package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      *LoginDetails
		want       string
		statusCode int
	}{
		{
			name:       "TestLoginSuccessful",
			method:     http.MethodGet,
			want:       "Login successful",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/login", bytes.NewBuffer([]byte("{\"username\": \"admin\", \"password\": \"admin\"}")))
			responseRecorder := httptest.NewRecorder()

			loginHandler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if responseRecorder.Body.String() != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
