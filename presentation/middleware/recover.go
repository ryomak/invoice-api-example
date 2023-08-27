package middleware

import (
	"errors"
	"github.com/ryomak/invoice-api-example/presentation/resource"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				resource.ErrorJson(
					w,
					errors.New("panic error"),
				)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
