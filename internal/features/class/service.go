package class

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"strings"
)

var ErrInvalidClassName = errors.New("invalid class name")

type ClassService struct {
	repo Repository
}

func NewClassService(repo Repository) *ClassService {
	return &ClassService{
		repo: repo,
	}
}

func (s *ClassService) Create(req CreateClassRequest) (Class, error) {
	name := strings.TrimSpace(req.Name)

	if name == "" {
		return Class{}, ErrInvalidClassName
	}

	joinCode, err := s.generateJoinCode()
	if err != nil {
		return Class{}, err
	}

	newClass := Class{
		Name:     name,
		JoinCode: joinCode,
	}

	return s.repo.Create(newClass)
}

func (s *ClassService) GetByID(id int) (Class, error) {
	return s.repo.GetByID(id)
}

func (s *ClassService) GetAll() ([]Class, error) {
	return s.repo.GetAll()
}

func (s *ClassService) Update(
	id int,
	req UpdateClassRequest,
) (Class, error) {
	currentClass, err := s.repo.GetByID(id)
	if err != nil {
		return Class{}, err
	}

	if req.Name != "" {
		name := strings.TrimSpace(req.Name)

		if name == "" {
			return Class{}, ErrInvalidClassName
		}

		currentClass.Name = name
	}

	if req.RequestJoinCode {
		joinCode, err := s.generateJoinCode()
		if err != nil {
			return Class{}, err
		}

		currentClass.JoinCode = joinCode
	}

	return s.repo.Update(id, currentClass)
}

func (s *ClassService) generateJoinCode() (string, error) {
	for {
		randomBytes := make([]byte, 5)

		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", err
		}

		value := base32.StdEncoding.
			WithPadding(base32.NoPadding).
			EncodeToString(randomBytes)

		joinCode := "CLS-" + value

		classes, err := s.repo.GetAll()
		if err != nil {
			return "", err
		}

		exists := false

		for _, currentClass := range classes {
			if currentClass.JoinCode == joinCode {
				exists = true
				break
			}
		}

		if !exists {
			return joinCode, nil
		}
	}
}
