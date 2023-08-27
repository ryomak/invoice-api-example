package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/domain/repository"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/converter"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type invoice struct {
	conn *db.Conn
}

func NewInvoice(
	conn *db.Conn,
) repository.Invoice {
	return &invoice{
		conn: conn,
	}
}

func (i *invoice) FindByCompanyIDAndFromTo(ctx context.Context, companyID uint64, from, to time.Time, offset, limit int) ([]*entity.Invoice, error) {
	mo, err := model.Invoices(
		model.InvoiceWhere.CompanyID.EQ(companyID),
		model.InvoiceWhere.DueAt.GTE(from),
		model.InvoiceWhere.DueAt.LTE(to),
		qm.Limit(limit),
		qm.Offset(offset),
	).All(ctx, i.conn)
	if err != nil {
		return nil, err
	}
	return converter.InvoicesToEntities(mo), nil
}

func (i *invoice) Create(ctx context.Context, invoice *entity.Invoice) error {
	mo := converter.InvoiceToModel(invoice)
	if err := mo.Insert(ctx, i.conn, boil.Infer()); err != nil {
		return err
	}
	invoice.ID = mo.ID
	return nil
}
