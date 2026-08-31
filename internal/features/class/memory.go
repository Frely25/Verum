package class

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

func (m *MemoryRepository) Create(class Class) (Class, error) {
	// Добавление в массив
	return Class{}, nil
}
func (m *MemoryRepository) GetAll() ([]Class, error) {
	return m.classes, nil
}

func (m *MemoryRepository) GetByID(id int) (Class, error) {
	// Логика for
	return Class{}, nil
}
