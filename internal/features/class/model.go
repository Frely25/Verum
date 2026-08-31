package class

// Model
type Class struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	JoinCode string `json:"join_code"`
}

// DTO - Data Transfer Object
type CreateClassRequest struct {
	Name string `json:"name"`
}
