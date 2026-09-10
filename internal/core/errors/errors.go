package apperrors

import "errors"

var (
	ErrClassNotFound    = errors.New("class not found")
	ErrInvalidClassName = errors.New("invalid class name")
)
