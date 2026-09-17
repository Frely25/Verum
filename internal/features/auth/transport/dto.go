package transport

import "github.com/Frely25/Verum/internal/core/domains"

type RegisterRequest struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
}

func userToResponse(user domains.User) UserResponse {

	return UserResponse{
		ID:          user.ID,
		Login:       user.Login,
		DisplayName: user.DisplayName,
	}
}
