package entity

import (
	"errors"
	"fmt"
	mtime "github.com/ryomak/invoice-api-example/pkg/time"
	"github.com/ryomak/invoice-api-example/pkg/unique"
	"time"
)

const (
	// FeeRatio 手数料率
	FeeRatio = 0.04
	// TaxRatio 消費税率
	TaxRatio = 0.1
)

type Invoice struct {
	ID              uint64
	RandID          string
	CompanyID       uint64        // 会社ID
	CompanyClientID uint64        // 取引先ID
	Status          InvoiceStatus // ステータス
	IssueAt         time.Time     // 発行日
	PaymentAmount   uint64        // 支払金額(取引先に支払う)
	BillingAmount   uint64        // 請求金額(会社に請求する金額)
	Fee             uint64        // 手数料
	FeeRatio        float64       // 手数料率
	Tax             uint64        // 消費税
	TaxRatio        float64       // 消費税率
	DueAt           time.Time     // 支払い期日
	CreatedAt       time.Time     // 作成日
	UpdatedAt       time.Time     // 更新日
}

// InvoiceStatus 請求書のステータス
type InvoiceStatus string

const (
	InvoiceStatusUnpaid     InvoiceStatus = "unpaid"     // 未払い
	InvoiceStatusPaid       InvoiceStatus = "paid"       // 支払い済み
	InvoiceStatusProcessing InvoiceStatus = "processing" // 処理中
	InvoiceStatusError      InvoiceStatus = "error"      // エラー
)

type invoiceBuilder struct {
	companyID       uint64
	companyClientID uint64
	issueAt         time.Time
	paymentAmount   uint64
	dueAt           time.Time
}

func NewInvoiceBuilder() *invoiceBuilder {
	return &invoiceBuilder{}
}

func (b *invoiceBuilder) CompanyID(companyID uint64) *invoiceBuilder {
	b.companyID = companyID
	return b
}

func (b *invoiceBuilder) CompanyClientID(companyClientID uint64) *invoiceBuilder {
	b.companyClientID = companyClientID
	return b
}

func (b *invoiceBuilder) IssueAt(issueAt time.Time) *invoiceBuilder {
	b.issueAt = issueAt
	return b
}

func (b *invoiceBuilder) PaymentAmount(paymentAmount uint64) *invoiceBuilder {
	b.paymentAmount = paymentAmount
	return b
}

func (b *invoiceBuilder) DueAt(dueAt time.Time) *invoiceBuilder {
	b.dueAt = dueAt
	return b
}

func (b *invoiceBuilder) CalcFee() float64 {
	return float64(b.paymentAmount) * FeeRatio
}

func (b *invoiceBuilder) CalcTax(fee float64) float64 {
	return fee * TaxRatio
}

func (b *invoiceBuilder) CalcBillingAmount(fee float64, tax float64) uint64 {
	return b.paymentAmount + uint64(fee) + uint64(tax)
}

func (b *invoiceBuilder) Build() (*Invoice, error) {
	if b.companyID == 0 {
		return nil, errors.New("companyID is required")
	}
	if b.companyClientID == 0 {
		return nil, errors.New("companyClientID is required")
	}
	if b.issueAt.IsZero() {
		return nil, errors.New("issueAt is required")
	}
	if b.paymentAmount == 0 {
		return nil, errors.New("paymentAmount is required")

	}
	if b.dueAt.IsZero() || b.dueAt.After(mtime.Now()) {
		return nil, errors.New("dueAt is required")
	}

	fee := b.CalcFee()
	tax := b.CalcTax(fee)
	billingAmount := b.CalcBillingAmount(fee, tax)

	return &Invoice{
		RandID:          fmt.Sprintf("inv-%s", unique.GenerateID()),
		CompanyID:       b.companyID,
		CompanyClientID: b.companyClientID,
		IssueAt:         b.issueAt,
		PaymentAmount:   b.paymentAmount,
		BillingAmount:   billingAmount,
		Fee:             uint64(fee),
		FeeRatio:        FeeRatio,
		Tax:             uint64(tax),
		TaxRatio:        TaxRatio,
		DueAt:           b.dueAt,
		Status:          InvoiceStatusUnpaid,
		CreatedAt:       mtime.Now(),
		UpdatedAt:       mtime.Now(),
	}, nil
}
