package class

import "errors"

type MemoryRepository struct {
	classes     []Class
	nextClassID int
}

func NewMemoryRepository() (*MemoryRepository, error) {
	return &MemoryRepository{
		classes:     make([]Class, 0),
		nextClassID: 1,
	}, nil
}

var ErrorNameNotFound error = errors.New("Name not found")

func (m *MemoryRepository) Create(class Class) (Class, error) {
	// Добавление в массив
	class.ID = m.nextClassID
	class.JoinCode = "ADC123"

	m.classes = append(m.classes, class)

	return class, nil
}
func (m *MemoryRepository) GetAll() ([]Class, error) {
	return m.classes, nil
}

func (m *MemoryRepository) GetByID(id int) (Class, error) {
	// Логика for
	for i := 0; i < len(m.classes); i++ {
		if id == m.classes[i].ID {
			return m.classes[i], nil
		}
	}
	return Class{}, ErrorNameNotFound
}
