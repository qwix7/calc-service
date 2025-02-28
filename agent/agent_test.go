package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Функция для фейковой отправки результата
func sendResult(taskID, result string) {
	// Этот фейковый сервер будет имитировать оркестратор
	http.HandleFunc("/internal/task/"+taskID, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// Отправляем результат
	resp, err := http.Post(server.URL+"/internal/task/"+taskID, "application/json", bytes.NewBufferString(result))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем, что ответ успешный
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestFetchTask(t *testing.T) {
	// Эмулируем ответ оркестратора
	task := Task{
		Expression: "2+2*2",
	}

	// Создаем фейковый сервер оркестратора
	tt := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(task)
	}))
	defer tt.Close()

	// Имитация получения задачи агентом
	got, err := fetchTask(tt.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if got.Expression != task.Expression {
		t.Fatalf("Expected task expression %s, got %s", task.Expression, got.Expression)
	}
}
