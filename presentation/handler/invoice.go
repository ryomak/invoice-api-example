package handler

import (
	"github.com/ryomak/invoice-api-example/application/request"
	"github.com/ryomak/invoice-api-example/application/usecase"
	"github.com/ryomak/invoice-api-example/pkg/logger"
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
		logger.Warningf(r.Context(), "InvoiceHandler.Get.Request: %v", err)
		resource.ErrorJson(w, err)
		return
	}

	res, err := h.invoiceUsecase.Get(r.Context(), req)
	if err != nil {
		logger.Errorf(r.Context(), "InvoiceHandler.Get: %v", err)
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
		logger.Warningf(r.Context(), "InvoiceHandler.Create.Request: %v", err)
		resource.ErrorJson(w, err)
		return
	}

	res, err := h.invoiceUsecase.Create(r.Context(), req)
	if err != nil {
		logger.Errorf(r.Context(), "InvoiceHandler.Create: %v", err)
		resource.ErrorJson(w, err)
		return
	}
	resource.JSON(w, res)
}
