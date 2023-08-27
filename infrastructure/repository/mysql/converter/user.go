package converter

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
	mtime "github.com/ryomak/invoice-api-example/pkg/time"
)

func UserToEntity(m *model.User) *entity.User {
	return &entity.User{
		ID:           m.ID,
		RandID:       m.RandID,
		Name:         m.Name,
		Mail:         m.Mail,
		CompanyID:    m.CompanyID,
		PasswordHash: m.PasswordHash,
		PasswordSalt: m.PasswordSalt,
	}
}

func UserToModel(e *entity.User) *model.User {
	return &model.User{
		ID:           e.ID,
		RandID:       e.RandID,
		Name:         e.Name,
		Mail:         e.Mail,
		CompanyID:    e.CompanyID,
		PasswordHash: e.PasswordHash,
		PasswordSalt: e.PasswordSalt,
		CreatedAt:    mtime.Now(),
		UpdatedAt:    mtime.Now(),
	}
}
