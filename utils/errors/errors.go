package errors

import (
	"errors"
	"net/http"
)

type GrpcError struct {
	Message string
	Code    uint16
	Err     string
}

func (e *GrpcError) Error() string {
	return e.Message
}

func NewError(message string) error {
	return errors.New(message)
}

func NewBadRequestError(message string) *GrpcError {
	return &GrpcError{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
	}
}

func NewNotFoundError(message string) *GrpcError {
	return &GrpcError{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string) *GrpcError {
	return &GrpcError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "internal_server_error",
	}
}
