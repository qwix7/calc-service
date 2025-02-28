package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskHandler_POST(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// Делаем POST-запрос с выражением
	reqBody := `{"expression": "2+3*3"}`
	resp, err := http.Post(server.URL+"/internal/task", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем, что ответ успешный
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Проверяем, что вернулся task_id
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if _, ok := result["task_id"]; !ok {
		t.Fatal("Expected task_id in response")
	}
}

func TestTaskHandler_GET(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// Эмулируем задачу
	taskID := "1"
	taskExpression := "2+2"

	// Сохраняем задачу в карте (это не рекомендуется для реальной реализации, это только для теста)
	tasks[taskID] = taskExpression

	// Делаем GET-запрос
	resp, err := http.Get(server.URL + "/internal/task/" + taskID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем успешность запроса
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Проверяем, что вернулся правильный ответ
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Expected valid JSON, got error: %v", err)
	}

	if result["expression"] != taskExpression {
		t.Fatalf("Expected expression %s, got %s", taskExpression, result["expression"])
	}
}
