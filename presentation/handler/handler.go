package handler

type Handler struct {
	Invoice InvoiceHandler
}

func New(
	invoice InvoiceHandler,
) *Handler {
	return &Handler{
		Invoice: invoice,
	}
}
