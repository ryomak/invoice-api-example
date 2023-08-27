package error

import (
	"database/sql"
	"errors"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func CodeByError(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return "0000"
}

func New(message string, code string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ErrBadRequest = New("bad request", "4001")
	ErrForbidden  = New("forbidden", "4002")
	ErrNoRows     = sql.ErrNoRows
)
