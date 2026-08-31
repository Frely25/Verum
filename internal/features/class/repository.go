package class

type Repository interface {
	Create(class Class) (Class, error)
	GetAll() ([]Class, error)
	GetByID(id int) (Class, error)
}
