package helpers

import "net/http"

type AppError struct {
	StatusCode int
	Message    string
}

func (se *AppError) Error() string {
	return se.Message
}

func NewInternalServerError() error {
	return &AppError{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
}

func NewNotFoundError() error {
	return &AppError{StatusCode: http.StatusNotFound, Message: "record not found"}
}
