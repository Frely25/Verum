package auth

import (
	"time"

	"github.com/Frely25/Verum/internal/core/domains"
)

type RegisterInput struct {
	Login       string
	Password    string
	DisplayName string
}

type LoginInput struct {
	Login    string
	Password string
}

type LoginResult struct {
	User      domains.User
	Token     string
	ExpiresAt time.Time
}
