package repository

import (
	"context"
	"sync"
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	users  []domains.User
	nextID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users:  make([]domains.User, 0),
		nextID: 1,
	}
}

func (r *MemoryRepository) Create(
	ctx context.Context,
	newUser domains.User,
) (domains.User, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, currentUser := range r.users {
		if currentUser.Login == newUser.Login {
			return domains.User{}, apperrors.ErrLoginAlreadyTaken
		}
	}

	now := time.Now()

	newUser.ID = r.nextID
	newUser.CreatedAt = now
	newUser.UpdatedAt = now

	r.nextID++

	r.users = append(r.users, newUser)

	return newUser, nil
}

func (r *MemoryRepository) GetByID(
	ctx context.Context,
	id int,
) (domains.User, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, currentUser := range r.users {
		if currentUser.ID == id {
			return currentUser, nil
		}
	}

	return domains.User{}, apperrors.ErrUserNotFound
}

func (r *MemoryRepository) GetByLogin(
	ctx context.Context,
	login string,
) (domains.User, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, currentUser := range r.users {
		if currentUser.Login == login {
			return currentUser, nil
		}
	}

	return domains.User{}, apperrors.ErrUserNotFound
}
