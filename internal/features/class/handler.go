package class

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) ClassesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateClass(w, r)
	case http.MethodGet:
		h.GetClasses(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetClassByID(w http.ResponseWriter, r *http.Request) {
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

	class, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "class not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
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
		Name: name,
	}

	class, err = h.repo.Create(class)
	if err != nil {
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func (h *Handler) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.repo.GetAll()
	if err != nil {
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "server is ok"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
