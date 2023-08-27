package converter

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
)

func CompanyToEntity(m *model.Company) *entity.Company {
	return &entity.Company{
		ID:          m.ID,
		RandID:      m.RandID,
		Name:        m.Name,
		PhoneNumber: m.PhoneNumber,
		PostalCode:  m.PostalCode,
		Address:     m.Address,
	}
}
