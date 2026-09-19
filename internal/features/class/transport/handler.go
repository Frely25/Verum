package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/core/tools"
	"github.com/Frely25/Verum/internal/features/class"
)

type Handler struct {
	ser Service
}

func NewHandler(ser Service) *Handler {
	return &Handler{
		ser: ser,
	}
}

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var reqCreate CreateClassRequest

	if err := json.NewDecoder(r.Body).Decode(&reqCreate); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	newClass, err := h.ser.Create(r.Context(), class.CreateInput{
		Name: reqCreate.Name,
	})

	switch {
	case errors.Is(err, apperrors.ErrInvalidClassName):
		http.Error(w, "invalid name", http.StatusUnprocessableEntity)
		return
	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tools.WriteJSON(w, http.StatusCreated, newClass)
}

func (h *Handler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		http.Error(w, "invalid class id", http.StatusBadRequest)
		return
	}

	currentClass, err := h.ser.GetByID(r.Context(), id)

	switch {
	case errors.Is(err, apperrors.ErrClassNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tools.WriteJSON(w, http.StatusOK, currentClass)
}

func (h *Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		http.Error(
			w,
			"invalid class id",
			http.StatusBadRequest,
		)
		return
	}

	var req UpdateClassRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	updatedClass, err := h.ser.Update(r.Context(), id, class.UpdateInput{
		Name:            req.Name,
		RequestJoinCode: req.RequestJoinCode,
	})

	switch {
	case errors.Is(err, apperrors.ErrClassNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	case errors.Is(err, apperrors.ErrInvalidClassName):
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return

	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tools.WriteJSON(w, http.StatusOK, updatedClass)
}
