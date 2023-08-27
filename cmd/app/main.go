package main

import (
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/env"
	"github.com/ryomak/invoice-api-example/presentation/router"
)

func main() {
	conn, err := db.New()
	if err != nil {
		panic(err)
	}
	r, err := router.New(conn)
	if err != nil {
		panic(err)
	}
	r.Routes()
	r.Run(env.GetCfg().Port)
}
