package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Template error:", err)
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Web interface running on :8081")
	http.ListenAndServe(":8081", nil)
}
