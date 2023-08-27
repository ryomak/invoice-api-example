package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
)

type Company interface {
	GetByUserID(ctx context.Context, userID uint64) (*entity.Company, error)
}
