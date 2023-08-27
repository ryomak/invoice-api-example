package entity

type User struct {
	ID        uint64
	RandID    string
	Name      string
	CompanyID uint64
	Mail      string

	PasswordHash string
	PasswordSalt string
}
