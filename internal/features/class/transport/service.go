package transport

import "github.com/Frely25/Verum/internal/core/domains"

type Service interface {
	Create(req CreateClassRequest) (domains.Class, error)
	GetByID(id int) (domains.Class, error)
	GetAll() ([]domains.Class, error)
	Update(id int, req UpdateClassRequest) (domains.Class, error)
}
