package repository

import (
	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type MemoryRepository struct {
	classes     []domains.Class
	nextClassID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		classes:     make([]domains.Class, 0),
		nextClassID: 1,
	}
}

func (m *MemoryRepository) Create(class domains.Class) (domains.Class, error) {
	// Добавление в массив
	class.ID = m.nextClassID
	m.nextClassID++

	m.classes = append(m.classes, class)

	return class, nil
}
func (m *MemoryRepository) GetAll() ([]domains.Class, error) {
	return m.classes, nil
}

func (m *MemoryRepository) GetByID(id int) (domains.Class, error) {
	// Логика for
	for i := 0; i < len(m.classes); i++ {
		if id == m.classes[i].ID {
			return m.classes[i], nil
		}
	}
	return domains.Class{}, apperrors.ErrClassNotFound
}

func (m *MemoryRepository) Delete(id int) (domains.Class, error) {
	for i := 0; i < len(m.classes); i++ {
		if m.classes[i].ID == id {
			deletedClass := m.classes[i]
			m.classes = append(
				m.classes[:i],
				m.classes[i+1:]...,
			)
			return deletedClass, nil
		}
	}
	return domains.Class{}, apperrors.ErrClassNotFound
}

func (m *MemoryRepository) Update(id int, classChanged domains.Class) (domains.Class, error) {
	for i := 0; i < len(m.classes); i++ {
		if m.classes[i].ID == id {
			m.classes[i] = classChanged
			return m.classes[i], nil
		}
	}
	return domains.Class{}, apperrors.ErrClassNotFound
}
