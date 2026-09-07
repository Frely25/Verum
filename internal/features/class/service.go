package class

type ServiceClass struct {
	repo Repository
}

func NewServiceClass(repo Repository) *ServiceClass {
	return &ServiceClass{
		repo: repo,
	}
}

func (ser *ServiceClass) Create(req CreateClassRequest) (Class, error) {
	return Class{}, nil
}

func (ser *ServiceClass) GetByID(id int) (Class, error) {
	return Class{}, nil
}

func (ser *ServiceClass) GetAll() ([]Class, error) {
	return []Class{}, nil
}

func (ser *ServiceClass) Update(id int, req UpdateClassRequest) (Class, error) {
	return Class{}, nil
}

func (ser *ServiceClass) generateJoinClass() string {
	joinClass := ""
	return joinClass
}
