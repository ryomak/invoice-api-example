package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/domain/repository"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/converter"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
)

type user struct {
	conn *db.Conn
}

func NewUser(
	conn *db.Conn,
) repository.User {
	return &user{
		conn: conn,
	}
}

func (u *user) GetByRandID(ctx context.Context, randID string) (*entity.User, error) {
	mo, err := model.Users(
		model.UserWhere.RandID.EQ(randID),
	).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}

	return converter.UserToEntity(mo), nil
}
