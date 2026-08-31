package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Model
type Class struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	JoinCode string `json:"join_code"`
}

// DTO - Data Transfer Object
type CreateClassRequest struct {
	Name string `json:"name"`
}

var classes = make([]Class, 0)
var nextClassID = 1

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthsHandler)
	mux.HandleFunc("/classes", classesHandler)
	mux.HandleFunc("GET /classes/{id}", getClassByID)

	log.Println("server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func classesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createClass(w, r)
	case http.MethodGet:
		getClass(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func getClassByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(
			w,
			"invalid class id",
			http.StatusBadRequest,
		)
		return
	}

	for i := 0; i < len(classes); i++ {
		if id == classes[i].ID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(classes[i])
			return
		}
	}

	http.Error(w, "class not found", http.StatusNotFound)
}

func createClass(w http.ResponseWriter, r *http.Request) {
	var req CreateClassRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(
			w,
			"invalid json",
			http.StatusBadRequest,
		)
		return
	}

	name := strings.TrimSpace(req.Name)

	if name == "" {
		http.Error(w, "invalid name", http.StatusBadRequest)
		return
	}
	class := Class{
		ID:       nextClassID,
		Name:     req.Name,
		JoinCode: "ABC123",
	}
	nextClassID++

	classes = append(classes, class)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func getClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(classes)
}

func healthsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "server is ok"})
}
