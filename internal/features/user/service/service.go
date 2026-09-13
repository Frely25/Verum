package service

import (
	"context"
	"strings"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, newUser domains.User) (domains.User, error) {
	newUser.Login = strings.TrimSpace(newUser.Login)
	newUser.DisplayName = strings.TrimSpace(newUser.DisplayName)

	if newUser.Login == "" ||
		newUser.DisplayName == "" ||
		newUser.PasswordHash == "" {

		return domains.User{}, apperrors.ErrInvalidUser
	}

	return s.repo.Create(ctx, newUser)
}

func (s *Service) GetByID(ctx context.Context, id int) (domains.User, error) {
	if id <= 0 {
		return domains.User{}, apperrors.ErrInvalidUser
	}

	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByLogin(ctx context.Context, login string) (domains.User, error) {
	login = strings.TrimSpace(login)

	if login == "" {
		return domains.User{}, apperrors.ErrInvalidUser
	}

	return s.repo.GetByLogin(ctx, login)
}
