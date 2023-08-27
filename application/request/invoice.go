package request

import (
	"fmt"
	"time"
)

type InvoiceGet struct {
	From  time.Time `json:"-" schema:"from"`
	To    time.Time `json:"-" schema:"to"`
	Limit int       `json:"-" schema:"limit"`
}

type InvoiceCreate struct {
	CompanyClientRandID string    `json:"companyClientRandId"`
	PaymentAmount       uint64    `json:"paymentAmount"`
	DueAt               time.Time `json:"dueAt"`
}

func (i *InvoiceCreate) Validate() error {
	if i.CompanyClientRandID == "" {
		return fmt.Errorf("companyClientRandId is required")
	}
	if i.PaymentAmount == 0 {
		return fmt.Errorf("paymentAmount is required")
	}
	if i.DueAt.IsZero() || i.DueAt.Before(time.Now()) {
		return fmt.Errorf("dueAt is required")
	}
	return nil
}
