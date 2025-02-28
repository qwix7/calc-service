package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskHandler_POST(t *testing.T) {
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	reqBody := `{"expression": "2+3*3"}`
	resp, err := http.Post(server.URL+"/internal/task", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if _, ok := result["task_id"]; !ok {
		t.Fatal("Expected task_id in response")
	}
}

func TestTaskHandler_GET(t *testing.T) {
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	taskID := "1"
	taskExpression := "2+2"

	tasks[taskID] = taskExpression

	resp, err := http.Get(server.URL + "/internal/task/" + taskID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if result["expression"] != taskExpression {
		t.Fatalf("Expected expression %s, got %s", taskExpression, result["expression"])
	}
}
