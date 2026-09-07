package class

type Service interface {
	Create(req CreateClassRequest) (Class, error)
	GetByID(id int) (Class, error)
	GetAll() ([]Class, error)
	Update(id int, req UpdateClassRequest) (Class, error)
}
