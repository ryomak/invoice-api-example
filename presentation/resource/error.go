package resource

import (
	"errors"
	merr "github.com/ryomak/invoice-api-example/pkg/error"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func StatusCodeByError(err error) int {
	var e *merr.Error
	if errors.As(err, &e) {
		switch e.Code {
		case merr.ErrForbidden.Code:
			return http.StatusForbidden

		}
	}
	if errors.Is(err, merr.ErrNoRows) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
