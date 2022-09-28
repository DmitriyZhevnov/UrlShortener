package apperror

import (
	"encoding/json"
)

var (
	internalServerError = newAppError("Internal Server error")
	badRequestError     = newAppError("Bad Request")
	errNotFound         = newAppError("Not found")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
}

func newAppError(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

func NewInternalServerError(developerMessage string) *AppError {
	err := internalServerError
	err.DeveloperMessage = developerMessage
	return err
}

func NewBadRequestError(developerMessage string) *AppError {
	err := badRequestError
	err.DeveloperMessage = developerMessage
	return err
}

func NewErrNotFound(developerMessage string) *AppError {
	err := errNotFound
	err.DeveloperMessage = developerMessage
	return err
}

func (a *AppError) Error() string {
	return a.Message
}

func (a *AppError) Unwrap() error {
	return a.Err
}

func (a *AppError) Marshal() []byte {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return marshal
}
