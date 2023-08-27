package middleware

type Middleware struct {
	AuthMiddleware AuthMiddleware
}

func New(
	authMiddleware AuthMiddleware,
) *Middleware {
	return &Middleware{
		AuthMiddleware: authMiddleware,
	}
}
