package transport

import (
	"context"
	"net/http"
)

type contextKey string

const currentUserIDKey contextKey = "current_user_id"

func (h *Handler) RequireAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx := context.WithValue(
			r.Context(),
			currentUserIDKey,
			currentUser.ID,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CurrentUserID(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(currentUserIDKey).(int)

	return userID, ok
}
