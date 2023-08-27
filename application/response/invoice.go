package response

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"time"
)

type InvoiceGet struct {
	Items []*Invoice `json:"items"`
}

func NewInvoiceGet(e []*entity.Invoice) *InvoiceGet {
	return &InvoiceGet{
		Items: NewInvoices(e),
	}
}

type InvoiceCreate struct {
	Invoice *Invoice `json:"invoice"`
}

func NewInvoiceCreate(e *entity.Invoice) *InvoiceCreate {
	return &InvoiceCreate{
		Invoice: NewInvoice(e),
	}
}

type Invoice struct {
	RandID        string               `json:"randId"`
	Status        entity.InvoiceStatus `json:"status"`
	IssueAt       time.Time            `json:"issueAt"`
	PaymentAmount uint64               `json:"paymentAmount"`
	BillingAmount uint64               `json:"billingAmount"`
	Fee           uint64               `json:"fee"`
	FeeRatio      float64              `json:"feeRatio"`
	Tax           uint64               `json:"tax"`
	TaxRatio      float64              `json:"taxRatio"`
	DueAt         time.Time            `json:"dueAt"`
}

func NewInvoice(e *entity.Invoice) *Invoice {
	return &Invoice{
		RandID:        e.RandID,
		Status:        e.Status,
		IssueAt:       e.IssueAt,
		PaymentAmount: e.PaymentAmount,
		BillingAmount: e.BillingAmount,
		Fee:           e.Fee,
		FeeRatio:      e.FeeRatio,
		Tax:           e.Tax,
		TaxRatio:      e.TaxRatio,
		DueAt:         e.DueAt,
	}
}

func NewInvoices(e []*entity.Invoice) []*Invoice {
	s := make([]*Invoice, 0, len(e))
	for _, v := range e {
		s = append(s, NewInvoice(v))
	}
	return s
}
