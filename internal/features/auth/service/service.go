package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/features/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users      UserService
	sessions   SessionRepository
	sessionTTL time.Duration
}

func NewService(users UserService, sessions SessionRepository) *Service {
	return &Service{
		users:    users,
		sessions: sessions,
	}
}

func (s *Service) Register(ctx context.Context, input auth.RegisterInput) (domains.User, error) {
	input.Login = strings.TrimSpace(input.Login)
	input.DisplayName = strings.TrimSpace(input.DisplayName)

	if input.Login == "" || input.Password == "" || input.DisplayName == "" {
		return domains.User{},
			apperrors.ErrInvalidAuthInput
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return domains.User{}, err
	}

	newUser := domains.User{
		Login:        input.Login,
		PasswordHash: string(passwordHash),
		DisplayName:  input.DisplayName,
	}

	return s.users.Create(ctx, newUser)
}

func (s *Service) Login(ctx context.Context, input auth.LoginInput) (auth.LoginResult, error) {
	input.Login = strings.TrimSpace(input.Login)

	if input.Login == "" || input.Password == "" {
		return auth.LoginResult{},
			apperrors.ErrInvalidCredentials
	}

	currentUser, err := s.users.GetByLogin(ctx, input.Login)

	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return auth.LoginResult{}, apperrors.ErrInvalidCredentials
		}

		return auth.LoginResult{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.PasswordHash), []byte(input.Password))

	if err != nil {
		return auth.LoginResult{},
			apperrors.ErrInvalidCredentials
	}

	token, expiresAt, err :=
		s.sessions.Create(
			ctx,
			currentUser.ID,
			s.sessionTTL,
		)

	if err != nil {
		return auth.LoginResult{}, err
	}

	return auth.LoginResult{
		User:      currentUser,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (domains.User, error) {
	token = strings.TrimSpace(token)

	if token == "" {
		return domains.User{}, apperrors.ErrUnauthorized
	}

	userID, err := s.sessions.Get(ctx, token)

	if err != nil {
		return domains.User{}, apperrors.ErrUnauthorized
	}

	currentUser, err := s.users.GetByID(ctx, userID)

	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return domains.User{}, apperrors.ErrUnauthorized
		}

		return domains.User{}, err
	}

	return currentUser, nil
}

func (s *Service) Logout(ctx context.Context, token string) error {

	if token == "" {
		return nil
	}

	return s.sessions.Delete(ctx, token)
}
