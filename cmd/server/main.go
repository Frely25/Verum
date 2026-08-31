package main

import (
	"log"
	"net/http"

	"github.com/Frely25/Verum/internal/features/class"
)

func main() {
	// var repo class.Repository
	// var service class.Service
	// var handler class.Handler

	repo := class.NewMemoryRepository()
	handler := class.NewHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/classes", handler.ClassesHandler)
	mux.HandleFunc("GET /classes/{id}", handler.GetClassByID)

	log.Println("server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
