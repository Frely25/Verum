package service

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"strings"

	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
	"github.com/Frely25/Verum/internal/features/class"
)

type ClassService struct {
	repo Repository
}

func NewClassService(repo Repository) *ClassService {
	return &ClassService{
		repo: repo,
	}
}

func (s *ClassService) Create(ctx context.Context, input class.CreateInput) (domains.Class, error) {
	name := strings.TrimSpace(input.Name)

	if name == "" {
		return domains.Class{}, apperrors.ErrInvalidClassName
	}

	for {
		joinCode, err := generateJoinCode()
		if err != nil {
			return domains.Class{}, err
		}

		newClass := domains.Class{
			Name:     name,
			JoinCode: joinCode,
		}

		createdClass, err := s.repo.Create(ctx, newClass)

		if errors.Is(err, apperrors.ErrJoinCodeAlreadyTaken) {
			continue
		}

		if err != nil {
			return domains.Class{}, err
		}

		return createdClass, nil
	}
}

func generateJoinCode() (string, error) {
	randomBytes := make([]byte, 5)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	value := base32.StdEncoding.
		WithPadding(base32.NoPadding).
		EncodeToString(randomBytes)

	return "CLS-" + value, nil
}

func (s *ClassService) GetByID(ctx context.Context, id int) (domains.Class, error) {
	if id <= 0 {
		return domains.Class{}, apperrors.ErrClassNotFound
	}

	return s.repo.GetByID(ctx, id)
}

func (s *ClassService) GetByJoinCode(ctx context.Context, joinCode string) (domains.Class, error) {
	joinCode = strings.TrimSpace(joinCode)

	if joinCode == "" {
		return domains.Class{}, apperrors.ErrClassNotFound
	}

	return s.repo.GetByJoinCode(ctx, joinCode)
}

func (s *ClassService) Update(ctx context.Context, id int, input class.UpdateInput) (domains.Class, error) {
	currentClass, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return domains.Class{}, err
	}

	changed := false

	if input.Name != "" {
		name := strings.TrimSpace(input.Name)

		if name == "" {
			return domains.Class{}, apperrors.ErrInvalidClassName
		}

		currentClass.Name = name
		changed = true
	}

	if !input.RequestJoinCode {
		if !changed {
			return currentClass, nil
		}

		return s.repo.Update(ctx, currentClass)
	}

	for {
		joinCode, err := generateJoinCode()
		if err != nil {
			return domains.Class{}, err
		}

		currentClass.JoinCode = joinCode

		updatedClass, err := s.repo.Update(ctx, currentClass)

		if errors.Is(err, apperrors.ErrJoinCodeAlreadyTaken) {
			continue
		}

		if err != nil {
			return domains.Class{}, err
		}

		return updatedClass, nil
	}
}
