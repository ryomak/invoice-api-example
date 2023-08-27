package entity

type Company struct {
	ID                 uint64
	RandID             string
	Name               string // 会社名
	RepresentativeName string // 代表者名
	PhoneNumber        string // 電話番号
	PostalCode         string // 郵便番号
	Address            string // 住所
}

type CompanyClient struct {
	ID                 uint64
	RandID             string
	CompanyID          uint64       // 会社ID
	Name               string       // 氏名
	RepresentativeName string       // 代表者名
	PhoneNumber        string       // 電話番号
	PostalCode         string       // 郵便番号
	Address            string       // 住所
	BankAccount        *BankAccount // 銀行口座
}

type BankAccount struct {
	ID         uint64
	Bank       *Bank       // 銀行
	BankBranch *BankBranch // 支店
	Number     string      // 口座番号
	HolderName string      // 口座名義
}
