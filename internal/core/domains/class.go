package domains

import "time"

type Class struct {
	ID        int
	Name      string
	JoinCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
