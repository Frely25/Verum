package class

type Repository interface {
	Create(class Class) (Class, error)
	GetAll() ([]Class, error)
	GetByID(id int) (Class, error)
	Delete(id int) (Class, error)
	Update(id int, classChanged Class) (Class, error)
}
