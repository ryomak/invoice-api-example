package converter

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
)

func CompanyClientToEntity(m *model.CompanyClient) *entity.CompanyClient {
	return &entity.CompanyClient{
		ID:                 m.ID,
		RandID:             m.RandID,
		Name:               m.Name,
		RepresentativeName: m.RepresentativeName,
		PhoneNumber:        m.PhoneNumber,
		PostalCode:         m.PostalCode,
		Address:            m.Address,
		BankAccount:        BankAccountToEntity(m.R.GetBankAccount()),
	}
}
