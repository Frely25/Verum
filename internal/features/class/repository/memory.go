package repository

import (
	"github.com/Frely25/Verum/internal/core/domains"
	apperrors "github.com/Frely25/Verum/internal/core/errors"
)

type MemoryRepository struct {
	classes     []classModel
	nextClassID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		classes:     make([]classModel, 0),
		nextClassID: 1,
	}
}

func (m *MemoryRepository) Create(class domains.Class) (domains.Class, error) {
	// Добавление в массив

	model := fromDomain(class)

	model.ID = m.nextClassID
	m.nextClassID++

	m.classes = append(m.classes, model)

	return toDomain(model), nil
}

func (m *MemoryRepository) GetAll() ([]domains.Class, error) {
	need_classes := make([]domains.Class, len(m.classes))

	for _, model := range m.classes {
		need_classes = append(need_classes, toDomain(model))
	}

	return need_classes, nil
}

func (m *MemoryRepository) GetByID(id int) (domains.Class, error) {
	// Логика for
	for i := 0; i < len(m.classes); i++ {
		if id == m.classes[i].ID {
			return toDomain(m.classes[i]), nil
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
			return toDomain(deletedClass), nil
		}
	}
	return domains.Class{}, apperrors.ErrClassNotFound
}

func (m *MemoryRepository) Update(id int, classChanged domains.Class) (domains.Class, error) {
	for i := 0; i < len(m.classes); i++ {
		if m.classes[i].ID == id {
			m.classes[i] = fromDomain(classChanged)
			return toDomain(m.classes[i]), nil
		}
	}
	return domains.Class{}, apperrors.ErrClassNotFound
}
