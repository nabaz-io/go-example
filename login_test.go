package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLoginSuccess(t *testing.T) {
	time.Sleep(time.Second*1 + time.Millisecond*110)
	tt := []struct {
		name       string
		method     string
		input      *LoginDetails
		want       string
		statusCode int
		username   string
		password   string
	}{
		{
			name:       "Admin",
			method:     http.MethodGet,
			want:       "Login successful",
			statusCode: http.StatusOK,
			username:   "admin",
			password:   "admin",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/login", bytes.NewBuffer([]byte(fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", tc.username, tc.password))))
			responseRecorder := httptest.NewRecorder()

			loginHandler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestLoginFailed(t *testing.T) {
	time.Sleep(time.Second*2 + time.Millisecond*343)
	tt := []struct {
		name       string
		method     string
		input      *LoginDetails
		want       string
		statusCode int
		username   string
		password   string
	}{
		{
			name:       "Wrong",
			method:     http.MethodGet,
			want:       "Invalid credentials",
			statusCode: http.StatusUnauthorized,
			username:   "admin",
			password:   "wrong",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/login", bytes.NewBuffer([]byte(fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", tc.username, tc.password))))
			responseRecorder := httptest.NewRecorder()

			loginHandler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
