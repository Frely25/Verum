package main

import (
	"Verum/internal/features/class"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", class.HealthsHandler)
	mux.HandleFunc("/classes", class.ClassesHandler)
	mux.HandleFunc("GET /classes/{id}", class.GetClassByID)

	log.Println("server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
