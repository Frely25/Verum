package domains

import "time"

type User struct {
	ID           int
	Login        string
	PasswordHash string
	DisplayName  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
