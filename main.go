package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateExpression(expression string) (float64, error) {
	valid := strings.ReplaceAll(expression, "+", "")
	valid = strings.ReplaceAll(valid, "-", "")
	valid = strings.ReplaceAll(valid, "*", "")
	valid = strings.ReplaceAll(valid, "/", "")
	valid = strings.ReplaceAll(valid, ".", "")
	if _, err := strconv.ParseFloat(valid, 64); err != nil {
		return 0, errors.New("Expression is not valid")
	}

	result, err := eval(expression)
	if err != nil {
		return 0, errors.New("Expression is not valid")
	}

	return result, nil
}

func eval(expression string) (float64, error) {
	return strconv.ParseFloat(expression, 64)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	result, err := calculateExpression(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: fmt.Sprintf("%.2f", result)})
}

func main() {
	http.HandleFunc("/api/v1/calculate", handler)
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
