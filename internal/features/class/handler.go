package class

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ClassesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateClass(w, r)
	case http.MethodGet:
		GetClass(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func GetClassByID(w http.ResponseWriter, r *http.Request) {
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

	// w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(classes[i])

	http.Error(w, "class not found", http.StatusNotFound)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
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

func GetClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(classes)
}

func HealthsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "server is ok"})
}
