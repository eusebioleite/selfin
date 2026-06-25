package errors

import "errors"

var (
	ErrUnauthorized = errors.New("User not authorized.")
	ErrNotFound     = errors.New("Resource not found.")
	ErrInternal     = errors.New("Internal server error.")
	ErrInvalidInput = errors.New("Invalid input.")
)

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
