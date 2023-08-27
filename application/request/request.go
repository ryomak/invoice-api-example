package request

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

func WithBody(req *http.Request, obj any) error {
	if err := json.NewDecoder(req.Body).Decode(obj); err != nil {
		return err
	}
	return nil
}

func WithQuery(req *http.Request, obj any) error {
	if err := schema.NewDecoder().Decode(obj, req.URL.Query()); err != nil {
		return err
	}
	return nil
}

type Validater interface {
	Validate() error
}

func NewParseByObject(req *http.Request, obj any, optionFunc ...func(req *http.Request, obj any) error) error {
	for _, fn := range optionFunc {
		if err := fn(req, obj); err != nil {
			return err
		}
	}
	if v, ok := obj.(Validater); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
