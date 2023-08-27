package middleware

import (
	"errors"
	"github.com/ryomak/invoice-api-example/pkg/logger"
	"github.com/ryomak/invoice-api-example/presentation/resource"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf(r.Context(), "Recover: %v", string(debug.Stack()))
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
