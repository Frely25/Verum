package service

import (
	"context"
	"strings"
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/features/auth/transport"
	"golang.org/x/crypto/bcrypt"
)

const sessionTTL = 24 * time.Hour

type Service struct {
	users    UserService
	sessions SessionRepository
}

func NewService(users UserService, sessions SessionRepository) *Service {
	return &Service{
		users:    users,
		sessions: sessions,
	}
}

func (s *Service) Register(ctx context.Context, req transport.RegisterRequest) (domains.User, error) {

	req.Login = strings.TrimSpace(req.Login)
	req.DisplayName = strings.TrimSpace(req.DisplayName)

	if req.Login == "" || req.Password == "" || req.DisplayName == "" {
		return domains.User{},
			apperrors.ErrInvalidAuthInput
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return domains.User{}, err
	}

	newUser := domains.User{
		Login:        req.Login,
		PasswordHash: string(passwordHash),
		DisplayName:  req.DisplayName,
	}

	return s.users.Create(ctx, newUser)
}
