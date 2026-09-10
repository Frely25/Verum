package service

import "github.com/Frely25/Verum/internal/core/domains"

type Repository interface {
	Create(class domains.Class) (domains.Class, error)
	GetAll() ([]domains.Class, error)
	GetByID(id int) (domains.Class, error)
	Delete(id int) (domains.Class, error)
	Update(id int, classChanged domains.Class) (domains.Class, error)
}
