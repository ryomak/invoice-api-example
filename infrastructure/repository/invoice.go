package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/domain/repository"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"time"
)

type invoice struct {
	conn *db.Conn
}

func NewInvoice(
	conn *db.Conn,
) repository.Invoice {
	return &invoice{
		conn: conn,
	}
}

func (i *invoice) FindByCompanyIDAndFromTo(ctx context.Context, companyID uint64, from, to time.Time, limit int) ([]*entity.Invoice, error) {

	return nil, nil
}
