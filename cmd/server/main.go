package main

import (
	"log"
	"net/http"

	"github.com/Frely25/Verum/internal/features/class/repository"
	"github.com/Frely25/Verum/internal/features/class/service"
	"github.com/Frely25/Verum/internal/features/class/transport"
)

func main() {
	repo := repository.NewMemoryRepository()
	service := service.NewClassService(repo)
	handler := transport.NewHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/classes", handler.ClassesHandler)
	mux.HandleFunc("GET /classes/{id}", handler.GetClassByID)
	mux.HandleFunc("PATCH /classes/{id}", handler.UpdateClass)

	log.Println("server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// PATCH 41.51.125.63:8080/classes/5
