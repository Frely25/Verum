package transport

import (
	"context"

	"github.com/Frely25/Verum/internal/core/domains"
	"github.com/Frely25/Verum/internal/features/auth"
)

type Service interface {
	Register(ctx context.Context, input auth.RegisterInput) (domains.User, error)
	Login(ctx context.Context, input auth.LoginInput) (auth.LoginResult, error)
	Authenticate(ctx context.Context, token string) (domains.User, error)
	Logout(ctx context.Context, token string) error
}
