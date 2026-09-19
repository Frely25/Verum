package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/features/auth"
)

const sessionCookieName = "verum_session"

type Handler struct {
	service      Service
	cookieSecure bool
}

func NewHandler(service Service, cookieSecure bool) *Handler {
	return &Handler{
		service:      service,
		cookieSecure: cookieSecure,
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(data)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.Register(
		r.Context(),
		auth.RegisterInput{
			Login:       req.Login,
			Password:    req.Password,
			DisplayName: req.DisplayName,
		},
	)

	switch {
	case errors.Is(err, apperrors.ErrInvalidAuthInput):
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	case errors.Is(err, apperrors.ErrLoginAlreadyTaken):
		http.Error(w, err.Error(), http.StatusConflict)
		return

	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, userToResponse(createdUser))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	result, err := h.service.Login(
		r.Context(),
		auth.LoginInput{
			Login:    req.Login,
			Password: req.Password,
		},
	)

	switch {
	case errors.Is(err, apperrors.ErrInvalidCredentials):
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return

	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    result.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cookieSecure,
		SameSite: http.SameSiteLaxMode,
		Expires:  result.ExpiresAt,
	},
	)

	writeJSON(
		w,
		http.StatusOK,
		userToResponse(result.User),
	)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookieName)

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	currentUser, err := h.service.Authenticate(r.Context(), cookie.Value)

	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	writeJSON(w, http.StatusOK, userToResponse(currentUser))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookieName)

	if err == nil {
		err = h.service.Logout(r.Context(), cookie.Value)

		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cookieSecure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1},
	)

	w.WriteHeader(http.StatusNoContent)
}
