package request

import "time"

type InvoiceGet struct {
	From  time.Time `json:"-" schema:"from"`
	To    time.Time `json:"-" schema:"to"`
	Limit int       `json:"-" schema:"limit"`
}

type InvoiceCreate struct {
}

func (i *InvoiceCreate) Validate() error {
	return nil
}
