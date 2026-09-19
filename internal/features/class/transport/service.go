package transport

import (
	"context"

	"github.com/Frely25/Verum/internal/core/domains"
	classfeature "github.com/Frely25/Verum/internal/features/class"
)

type Service interface {
	Create(ctx context.Context, input classfeature.CreateInput) (domains.Class, error)
	GetByID(ctx context.Context, id int) (domains.Class, error)
	Update(ctx context.Context, id int, input classfeature.UpdateInput) (domains.Class, error)
}
