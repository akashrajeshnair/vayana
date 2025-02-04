package errors

import "fmt"

type ErrorCode int

const (
	// Common error codes starting from 1000
	ErrInvalidInput ErrorCode = 1000 + iota
	ErrUnauthorized
	ErrForbidden
	ErrNotFound
	ErrInternalServer
	ErrDatabaseOperation
	ErrInvalidToken
)

type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("error code: %d, message: %s", e.Code, e.Message)
}

// Error constructors
func NewInvalidInputError(msg string, err error) *AppError {
	return &AppError{
		Code:    ErrInvalidInput,
		Message: msg,
		Err:     err,
	}
}

func NewUnauthorizedError(msg string) *AppError {
	return &AppError{
		Code:    ErrUnauthorized,
		Message: msg,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code:    ErrNotFound,
		Message: msg,
	}
}

func NewDatabaseError(msg string, err error) *AppError {
	return &AppError{
		Code:    ErrDatabaseOperation,
		Message: msg,
		Err:     err,
	}
}
