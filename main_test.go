package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestMain(t *testing.T) {
	go main()

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/calculate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.DefaultServeMux
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}
}
