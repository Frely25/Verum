package service

import (
	"context"

	"github.com/Frely25/Verum/internal/core/domains"
)

type Repository interface {
	Create(ctx context.Context, newClass domains.Class) (domains.Class, error)
	GetByID(ctx context.Context, id int) (domains.Class, error)
	GetByJoinCode(ctx context.Context, joinCode string) (domains.Class, error)
	GetAll(ctx context.Context) ([]domains.Class, error)
	Update(ctx context.Context, classChanged domains.Class) (domains.Class, error)
}
