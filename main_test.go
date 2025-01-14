package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid expression",
			requestBody:    `{"expression": "2+2*2"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":"6.00"}`,
		},
		{
			name:           "Invalid expression with letters",
			requestBody:    `{"expression": "2+abc"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error":"Expression is not valid"}`,
		},
		{
			name:           "Empty expression",
			requestBody:    `{"expression": ""}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error":"Expression is not valid"}`,
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"expr": "2+2"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name:           "Unsupported HTTP method",
			requestBody:    "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(test.requestBody))
			req.Header.Set("Content-Type", "application/json")

			if test.name == "Unsupported HTTP method" {
				req.Method = http.MethodGet
			}

			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()
			body := w.Body.String()

			if resp.StatusCode != test.expectedStatus {
				t.Errorf("expected status %d, got %d", test.expectedStatus, resp.StatusCode)
			}

			if test.expectedBody != "" && body != test.expectedBody {
				t.Errorf("expected body %s, got %s", test.expectedBody, body)
			}
		})
	}
}
