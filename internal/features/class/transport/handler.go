package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type Handler struct {
	ser Service
}

func NewHandler(ser Service) *Handler {
	return &Handler{
		ser: ser,
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

	class, err := h.ser.GetByID(id)

	if errors.Is(err, apperrors.ErrClassNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	class, err := h.ser.Create(req)

	if errors.Is(err, apperrors.ErrInvalidClassName) {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if err != nil {
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func (h *Handler) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.ser.GetAll()
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

func (h *Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	var req UpdateClassRequest

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

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		return
	}
	changedClass, err := h.ser.Update(id, req)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(changedClass)
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "server is ok"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
