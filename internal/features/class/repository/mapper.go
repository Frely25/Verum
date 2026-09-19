package repository

import "github.com/Frely25/Verum/internal/core/domains"

func toDomain(model classModel) domains.Class {
	return domains.Class{
		ID:        model.ID,
		Name:      model.Name,
		JoinCode:  model.JoinCode,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func fromDomain(domain domains.Class) classModel {
	return classModel{
		ID:        domain.ID,
		Name:      domain.Name,
		JoinCode:  domain.JoinCode,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
