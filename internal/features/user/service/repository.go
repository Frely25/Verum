package service

import (
	"context"

	"github.com/Frely25/Verum/internal/core/domains"
)

type Repository interface {
	Create(ctx context.Context, newUser domains.User) (domains.User, error)
	GetByID(ctx context.Context, id int) (domains.User, error)
	GetByLogin(ctx context.Context, login string) (domains.User, error)
}
