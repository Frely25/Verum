package repository

import "time"

type classModel struct {
	ID        int
	Name      string
	JoinCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
