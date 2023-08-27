package converter

import (
	"github.com/ryomak/invoice-api-example/domain/entity"
	"github.com/ryomak/invoice-api-example/infrastructure/repository/mysql/model"
)

func BankToEntity(m *model.Bank) *entity.Bank {
	return &entity.Bank{
		ID:     m.ID,
		RandID: m.RandID,
		Name:   m.Name,
	}
}

func BankBranchToEntity(m *model.BankBranch) *entity.BankBranch {
	return &entity.BankBranch{
		ID:   m.ID,
		Name: m.Name,
	}
}

func BankAccountToEntity(m *model.BankAccount) *entity.BankAccount {
	return &entity.BankAccount{
		ID:         m.ID,
		Bank:       BankToEntity(m.R.GetBank()),
		BankBranch: BankBranchToEntity(m.R.GetBranch()),
		Number:     m.Number,
		HolderName: m.HolderName,
	}
}
