package apperror

import (
	"encoding/json"
)

var (
	internalServerError = newAppError("Internal Server error")
	badRequestError     = newAppError("Bad Request")
	errNotFound         = newAppError("Not found")
)

type appError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
}

func newAppError(message string) *appError {
	return &appError{
		Message: message,
	}
}

func NewInternalServerError(developerMessage string) *appError {
	err := internalServerError
	err.DeveloperMessage = developerMessage
	return err
}

func NewBadRequestError(developerMessage string) *appError {
	err := badRequestError
	err.DeveloperMessage = developerMessage
	return err
}

func NewErrNotFound(developerMessage string) *appError {
	err := errNotFound
	err.DeveloperMessage = developerMessage
	return err
}

func (a *appError) Error() string {
	return a.Message
}

func (a *appError) Unwrap() error {
	return a.Err
}

func (a *appError) Marshal() []byte {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return marshal
}
