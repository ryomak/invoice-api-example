package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/ryomak/invoice-api-example/di"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/presentation/handler"
	"github.com/ryomak/invoice-api-example/presentation/middleware"
	"log"
	"net/http"
)

type Router struct {
	chi.Router
	middleware *middleware.Middleware
	handler    *handler.Handler
}

func New(db *db.Conn) (*Router, error) {
	r := chi.NewRouter()
	m, err := di.Middlewares(db)
	if err != nil {
		return nil, err
	}
	h, err := di.Handlers(db)
	if err != nil {
		return nil, err
	}
	return &Router{Router: r, middleware: m, handler: h}, nil
}

func (router *Router) Routes() {
	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.Recover)
		r.Use(middleware.AccessLog)
		// 認証
		r.Group(func(r chi.Router) {
			r.Use(router.middleware.AuthMiddleware.Auth)
			r.Route("/invoices", func(r chi.Router) {
				r.Get("/", router.handler.Invoice.Get)
				r.Post("/", router.handler.Invoice.Create)
			})
		})

	})
}

func (router *Router) Run(port int) {
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
