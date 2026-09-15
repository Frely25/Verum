package apperrors

import "errors"

var (
	// Class errors
	ErrClassNotFound    = errors.New("class not found")
	ErrInvalidClassName = errors.New("invalid class name")

	// User errors
	ErrUserNotFound      = errors.New("user not found")
	ErrLoginAlreadyTaken = errors.New("login already taken")
	ErrInvalidUser       = errors.New("invalid user")

	// Auth errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrSessionNotFound    = errors.New("session not found")
	ErrInvalidAuthInput   = errors.New("invalid auth input")
)
