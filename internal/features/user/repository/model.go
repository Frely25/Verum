package repository

import (
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
)

type userModel struct {
	ID           int
	Login        string
	PasswordHash string
	DisplayName  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m userModel) toDomain() domains.User {
	return domains.User{
		ID:           m.ID,
		Login:        m.Login,
		PasswordHash: m.PasswordHash,
		DisplayName:  m.DisplayName,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}
