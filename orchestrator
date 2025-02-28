package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var tasks = make(map[string]string)

type Task struct {
	Expression string `json:"expression"`
}

type Result struct {
	Result string `json:"result"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		mu.Lock()
		taskID := fmt.Sprintf("%d", len(tasks)+1)
		tasks[taskID] = task.Expression
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"task_id": taskID})

	case http.MethodGet:
		taskID := r.URL.Path[len("/internal/task/"):]
		mu.Lock()
		task, exists := tasks[taskID]
		mu.Unlock()
		if !exists {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"expression": task})

	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/internal/task", taskHandler)
	log.Println("Orchestrator running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
