package middleware

import (
	myContext "github.com/ryomak/invoice-api-example/domain/context"
	"github.com/ryomak/invoice-api-example/domain/repository"
	"github.com/ryomak/invoice-api-example/infrastructure/env"
	"github.com/ryomak/invoice-api-example/pkg/logger"
	"github.com/ryomak/invoice-api-example/presentation/resource"
	"net/http"
)

type AuthMiddleware interface {
	Auth(next http.Handler) http.Handler
}

type authMiddleware struct {
	userRepository repository.User
}

func NewAuthMiddleware(
	userRepository repository.User,
) AuthMiddleware {
	return &authMiddleware{
		userRepository: userRepository,
	}
}

func (m *authMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO 認証処理を挟む
		// デバッグ
		if env.GetCfg().IsLocal() {
			userID := r.Header.Get("X-Debug-Id")
			user, err := m.userRepository.GetByRandID(r.Context(), userID)
			if err != nil {
				logger.Warningf(r.Context(), "AuthMiddleware.Auth: %v", err)
				resource.ErrorJson(w, err)
				return
			}
			r = r.WithContext(myContext.WithAuth(r.Context(), &myContext.Auth{
				User: user,
			}))

			// loggerにも詰める
			r = r.WithContext(logger.WithUser(r.Context(), user.RandID))

		} else {
			panic("implement me")
		}

		next.ServeHTTP(w, r)
	})
}
