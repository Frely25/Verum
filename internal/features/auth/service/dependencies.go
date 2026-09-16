package service

import (
	"context"
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
)

type UserService interface {
	Create(ctx context.Context, user domains.User) (domains.User, error)
	GetByID(ctx context.Context, id int) (domains.User, error)
	GetByLogin(ctx context.Context, login string) (domains.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, userID int, ttl time.Duration) (string, time.Time, error)
	Get(ctx context.Context, token string) (int, error)
	Delete(ctx context.Context, token string) error
}
