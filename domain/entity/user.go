package entity

type User struct {
	ID        uint64
	RandID    string
	Name      string // 氏名
	CompanyID uint64 // 会社ID
	Mail      string // メールアドレス

	PasswordHash string // パスワードハッシュ
	PasswordSalt string // パスワードソルト
}
