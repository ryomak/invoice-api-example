package handler

import (
	"context"
	"github.com/ryomak/invoice-api-example/application/request"
	"github.com/ryomak/invoice-api-example/application/usecase"
	"github.com/ryomak/invoice-api-example/presentation/resource"
	"net/http"
)

type InvoiceHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type invoiceHandler struct {
	invoiceUsecase usecase.Invoice
}

func NewInvoiceHandler(invoiceUsecase usecase.Invoice) InvoiceHandler {
	return &invoiceHandler{
		invoiceUsecase: invoiceUsecase,
	}
}

func (h *invoiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	var req = new(request.InvoiceGet)
	if err := request.NewParseByObject(
		r,
		req,
		request.WithQuery,
	); err != nil {
		resource.ErrorJson(w, err)
		return
	}

	res, err := h.invoiceUsecase.Get(context.Background(), req)
	if err != nil {
		resource.ErrorJson(w, err)
		return
	}

	resource.JSON(w, res)
}

func (h *invoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req = new(request.InvoiceCreate)
	if err := request.NewParseByObject(
		r,
		req,
		request.WithBody,
	); err != nil {
		resource.ErrorJson(w, err)
		return
	}

	res, err := h.invoiceUsecase.Create(context.Background(), req)
	if err != nil {
		resource.ErrorJson(w, err)
		return
	}
	resource.JSON(w, res)
}
