package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"time"
)

type Invoice interface {
	FindByCompanyIDAndFromTo(ctx context.Context, companyID uint64, from, to time.Time, limit int) ([]*entity.Invoice, error)
	Create(ctx context.Context, invoice *entity.Invoice) error
}
