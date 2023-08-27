package converter

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
)

const (
	InvoiceRatioDigit = 1000
)

func InvoiceToEntity(m *model.Invoice) *entity.Invoice {
	return &entity.Invoice{
		ID:              m.ID,
		RandID:          m.RandID,
		CompanyID:       m.CompanyID,
		CompanyClientID: m.CompanyClientID,
		Status:          entity.InvoiceStatus(m.Status),
		IssueAt:         m.IssueAt,
		PaymentAmount:   m.PaymentAmount,
		BillingAmount:   m.BillingAmount,
		Fee:             m.Fee,
		FeeRatio:        float64(m.FeeRatio / InvoiceRatioDigit),
		Tax:             m.Tax,
		TaxRatio:        float64(m.TaxRatio / InvoiceRatioDigit),
		DueAt:           m.DueAt,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func InvoiceToModel(e *entity.Invoice) *model.Invoice {
	return &model.Invoice{
		ID:              e.ID,
		RandID:          e.RandID,
		CompanyID:       e.CompanyID,
		CompanyClientID: e.CompanyClientID,
		Status:          string(e.Status),
		IssueAt:         e.IssueAt,
		PaymentAmount:   e.PaymentAmount,
		BillingAmount:   e.BillingAmount,
		Fee:             e.Fee,
		FeeRatio:        uint(e.FeeRatio * InvoiceRatioDigit),
		Tax:             e.Tax,
		TaxRatio:        uint(e.TaxRatio * InvoiceRatioDigit),
		DueAt:           e.DueAt,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}

func InvoicesToEntities(m []*model.Invoice) []*entity.Invoice {
	s := make([]*entity.Invoice, 0, len(m))
	for _, v := range m {
		s = append(s, InvoiceToEntity(v))
	}
	return s
}
