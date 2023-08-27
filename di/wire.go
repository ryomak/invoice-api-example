//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryomak/invoice-api-example/application/usecase"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/repository"
	"github.com/ryomak/invoice-api-example/presentation/handler"
	"github.com/ryomak/invoice-api-example/presentation/middleware"
)

func Handlers(
	conn *db.Conn,
) (*handler.Handler, error) {
	wire.Build(
		handler.New,
		handler.NewInvoiceHandler,
		usecase.NewInvoice,
	)

	return nil, nil
}

func Middlewares(
	conn *db.Conn,
) (*middleware.Middleware, error) {
	wire.Build(
		middleware.New,
		middleware.NewAuthMiddleware,
		repository.NewUser,
	)
	return nil, nil
}
