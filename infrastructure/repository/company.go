package repository

import (
	"context"
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/domain/repository"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/converter"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type company struct {
	conn *db.Conn
}

func NewCompany(
	conn *db.Conn,
) repository.Company {
	return &company{
		conn: conn,
	}
}

func (c *company) GetByUserID(ctx context.Context, userID uint64) (*entity.Company, error) {
	mo, err := model.Companies(
		InnerJoin(model.TableNames.User, model.TableNames.Company, model.UserColumns.CompanyID, model.CompanyColumns.ID),
		model.UserWhere.ID.EQ(userID),
	).One(ctx, c.conn)
	if err != nil {
		return nil, err
	}

	return converter.CompanyToEntity(mo), nil
}

func (c *company) GetClientByRandID(ctx context.Context, randID string) (*entity.CompanyClient, error) {
	mo, err := model.CompanyClients(
		model.CompanyClientWhere.RandID.EQ(randID),
		qm.Load(model.CompanyClientRels.BankAccount),
		qm.Load(qm.Rels(
			model.CompanyClientRels.BankAccount,
			model.BankAccountRels.Branch,
		)),
		qm.Load(qm.Rels(
			model.CompanyClientRels.BankAccount,
			model.BankAccountRels.Branch,
			model.BankBranchRels.Bank,
		)),
	).One(ctx, c.conn)
	if err != nil {
		return nil, err
	}

	return converter.CompanyClientToEntity(mo), nil
}
