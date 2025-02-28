package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// Делаем GET-запрос
	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем успешность запроса
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
