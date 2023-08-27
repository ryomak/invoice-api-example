package resource

import (
	"encoding/json"
	merr "github.com/ryomak/invoice-api-example/pkg/error"
	"net/http"
)

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func JSON(w http.ResponseWriter, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func ErrorJson(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCodeByError(err))
	_ = json.NewEncoder(w).Encode(&Error{
		Message: err.Error(),
		Code:    merr.CodeByError(err),
	})
}
