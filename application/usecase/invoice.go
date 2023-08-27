package usecase

import (
	"context"
	"github.com/ryomak/invoice-api-example/application/request"
	"github.com/ryomak/invoice-api-example/application/response"
)

type Invoice interface {
	Get(ctx context.Context, req *request.InvoiceGet) (*response.InvoiceGet, error)
	Create(ctx context.Context, req *request.InvoiceCreate) (*response.InvoiceCreate, error)
}

type invoice struct {
}

func NewInvoice() Invoice {
	return &invoice{}
}

func (i *invoice) Get(ctx context.Context, req *request.InvoiceGet) (*response.InvoiceGet, error) {
	return nil, nil
}

func (i *invoice) Create(ctx context.Context, req *request.InvoiceCreate) (*response.InvoiceCreate, error) {
	return nil, nil
}
