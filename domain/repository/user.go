package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
)

type User interface {
	GetByRandID(ctx context.Context, randID string) (*entity.User, error)
}
