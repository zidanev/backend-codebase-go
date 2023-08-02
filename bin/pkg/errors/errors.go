package errors

import (
	"net/http"
)

type ErrorString struct {
	code    int
	message string
}

func (e ErrorString) Code() int {
	return e.code
}

func (e ErrorString) Error() string {
	return e.message
}

func (e ErrorString) Message() string {
	return e.message
}

// BadRequest will throw if the given request-body or params is not valid
func BadRequest(msg string) error {
	return &ErrorString{
		code:    http.StatusBadRequest,
		message: msg,
	}
}

// NotFound will throw if the requested item is not exists
func NotFound(msg string) error {
	return &ErrorString{
		code:    http.StatusNotFound,
		message: msg,
	}
}

// Conflict will throw if the current action already exists
func Conflict(msg string) error {
	return &ErrorString{
		code:    http.StatusConflict,
		message: msg,
	}
}

// InternalServerError will throw if any the Internal Server Error happen,
// Database, Third Party etc.
func InternalServerError(msg string) error {
	return &ErrorString{
		code:    http.StatusInternalServerError,
		message: msg,
	}
}

func UnauthorizedError(msg string) error {
	return &ErrorString{
		code:    http.StatusUnauthorized,
		message: msg,
	}
}

func ForbiddenError(msg string) error {
	return &ErrorString{
		code:    http.StatusForbidden,
		message: msg,
	}
}
