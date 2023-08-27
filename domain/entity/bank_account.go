package entity

type BankAccount struct {
	ID         uint64
	Bank       *Bank       // 銀行名
	BankBranch *BankBranch // 支店名
	Number     string      // 口座番号
	HolderName string      // 口座名義
}

type Bank struct {
	ID   uint64
	Name string // 銀行名
}

type BankBranch struct {
	ID     uint64
	BankID uint64
	Name   string // 支店名
}
