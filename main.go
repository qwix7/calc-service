package main

import (
	"fmt"
	"log"
	"net/http"
	"calc-service/orchestrator"
	"calc-service/agent"
	"calc-service/web"
)

func main() {
	go orchestrator.StartOrchestrator()
	agent.StartAgents()
	http.HandleFunc("/api/v1/calculate", web.CalculateHandler)
	
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
