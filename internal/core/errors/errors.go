package apperrors

import "errors"

var (
	// Class
	ErrClassNotFound    = errors.New("class not found")
	ErrInvalidClassName = errors.New("invalid class name")

	// User
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidUser       = errors.New("invalid user")
	ErrLoginAlreadyTaken = errors.New("login already taken")
)
