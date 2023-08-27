package usecase

import (
	"context"
	"fmt"
	"github.com/ryomak/invoice-api-example/application/request"
	"github.com/ryomak/invoice-api-example/application/response"
	myContext "github.com/ryomak/invoice-api-example/domain/context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/domain/repository"
	mtime "github.com/ryomak/invoice-api-example/pkg/time"
)

type Invoice interface {
	Get(ctx context.Context, req *request.InvoiceGet) (*response.InvoiceGet, error)
	Create(ctx context.Context, req *request.InvoiceCreate) (*response.InvoiceCreate, error)
}

type invoice struct {
	companyRepository repository.Company
	invoiceRepository repository.Invoice
}

func NewInvoice(
	companyRepository repository.Company,
	invoiceRepository repository.Invoice,
) Invoice {
	return &invoice{
		companyRepository: companyRepository,
		invoiceRepository: invoiceRepository,
	}
}

func (i *invoice) Get(ctx context.Context, req *request.InvoiceGet) (*response.InvoiceGet, error) {
	auth := myContext.GetAuth(ctx)
	company, err := i.companyRepository.GetByUserID(ctx, auth.User.ID)
	if err != nil {
		return nil, fmt.Errorf("companyRepository.GetByUserID: %w", err)
	}
	invoices, err := i.invoiceRepository.FindByCompanyIDAndFromTo(ctx, company.ID, req.From, req.To, req.Limit)
	if err != nil {
		return nil, fmt.Errorf("invoiceRepository.FindByCompanyIDAndFromTo: %w", err)
	}

	return response.NewInvoiceGet(invoices), nil
}

func (i *invoice) Create(ctx context.Context, req *request.InvoiceCreate) (*response.InvoiceCreate, error) {
	auth := myContext.GetAuth(ctx)
	company, err := i.companyRepository.GetByUserID(ctx, auth.User.ID)
	if err != nil {
		return nil, fmt.Errorf("companyRepository.GetByUserID: %w", err)
	}
	client, err := i.companyRepository.GetClientByRandID(ctx, req.CompanyClientRandID)
	if err != nil {
		return nil, fmt.Errorf("companyRepository.GetClientByRandID: %w", err)
	}
	if company.ID != client.CompanyID {
		return nil, fmt.Errorf("companyRepository.GetClientByRandID: %w", err)
	}

	invoiceEnt, err := entity.NewInvoiceBuilder().
		CompanyID(company.ID).
		CompanyClientID(client.ID).
		IssueAt(mtime.Now()).
		PaymentAmount(req.PaymentAmount).
		DueAt(req.DueAt).
		Build()
	if err != nil {
		return nil, fmt.Errorf("entity.NewInvoiceBuilder: %w", err)
	}

	if err := i.invoiceRepository.Create(ctx, invoiceEnt); err != nil {
		return nil, fmt.Errorf("invoiceRepository.Create: %w", err)
	}

	return response.NewInvoiceCreate(invoiceEnt), nil
}
