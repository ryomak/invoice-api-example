package di

import "github.com/ryomak/invoice-api-example/infrastructure/client/db"

type ExternalClient struct {
	db *db.Conn
}

func NewExternalClient(db *db.Conn) *ExternalClient {
	return &ExternalClient{db: db}
}
