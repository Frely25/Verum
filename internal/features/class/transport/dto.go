package transport

// DTO - Data Transfer Object
type CreateClassRequest struct {
	Name string `json:"name"`
}

type UpdateClassRequest struct {
	Name            string `json:"name"`
	RequestJoinCode bool   `json:"request_join_code"`
}
