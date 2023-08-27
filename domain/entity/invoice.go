package entity

import "time"

type Invoice struct {
	ID              uint64
	RandID          string
	CompanyID       uint64        // 会社ID
	CompanyClientID uint64        // 取引先ID
	Status          InvoiceStatus // ステータス
	IssueAt         time.Time     // 発行日
	Amount          uint64        // 金額
	Fee             uint          // 手数料
	FeeRatio        float64       // 手数料率
	Tax             uint64        // 消費税
	TaxRatio        float64       // 消費税率
	DueAt           time.Time     // 支払い期日
	CreatedAt       time.Time     // 作成日
	UpdatedAt       time.Time     // 更新日
}

type InvoiceStatus string

const (
	InvoiceStatusUnpaid     InvoiceStatus = "unpaid"     // 未払い
	InvoiceStatusPaid       InvoiceStatus = "paid"       // 支払い済み
	InvoiceStatusProcessing InvoiceStatus = "processing" // 処理中
	InvoiceStatusError      InvoiceStatus = "error"      // エラー
)
