package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type session struct {
	UserID    int
	ExpiresAt time.Time
}

type MemoryRepository struct {
	mu       sync.RWMutex
	sessions map[string]session
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		sessions: make(map[string]session),
	}
}

func (r *MemoryRepository) Create(ctx context.Context, userID int, ttl time.Duration) (string, time.Time, error) {

	// Пустые байты
	randomBytes := make([]byte, 32)

	// Читает байты и выдает случайные безопасные числа от ОС
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", time.Time{}, err
	}

	// EncodeToString - перевод байтов в символы; RawURLEncoding - убирает (=)
	token := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Время окончания; ttl - time to live
	expiresAt := time.Now().Add(ttl)

	r.mu.Lock()

	r.sessions[token] = session{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}

	r.mu.Unlock()

	return token, expiresAt, nil
}

func (r *MemoryRepository) Get(ctx context.Context, token string) (int, error) {

	r.mu.RLock()
	currentSession, exists := r.sessions[token]
	r.mu.RUnlock()

	if !exists {
		return 0, apperrors.ErrSessionNotFound
	}

	if time.Now().After(currentSession.ExpiresAt) {
		_ = r.Delete(ctx, token)

		return 0, apperrors.ErrSessionNotFound
	}

	return currentSession.UserID, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, token string) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.sessions, token)

	return nil
}
