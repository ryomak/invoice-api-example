package entity

import (
	cmp "github.com/google/go-cmp/cmp"
	mtime "github.com/ryomak/invoice-api-example/pkg/time"
	"github.com/ryomak/invoice-api-example/pkg/unique"
	"testing"
	"time"
)

func TestInvoiceBuilder_Build(t *testing.T) {
	now := mtime.Now()
	mtime.SetFakeNow(now)
	unique.SetFakeGenerateID("id")
	type in struct {
		companyID       uint64
		companyClientID uint64
		issueAt         time.Time
		paymentAmount   uint64
		dueAt           time.Time
	}
	tests := []struct {
		name    string
		in      in
		want    *Invoice
		wantErr bool
	}{
		{
			name: "case1: normal",
			in: in{
				companyID:       1,
				companyClientID: 1,
				issueAt:         mtime.Now(),
				paymentAmount:   10000,
				dueAt:           mtime.Now().AddDate(0, 1, 0),
			},
			want: &Invoice{
				ID:              0,
				RandID:          "inv-id",
				CompanyID:       1,
				CompanyClientID: 1,
				IssueAt:         now,
				PaymentAmount:   10000,
				BillingAmount:   10440,
				Fee:             400,
				FeeRatio:        0.04,
				Tax:             40,
				TaxRatio:        0.1,
				DueAt:           now.AddDate(0, 1, 0),
				Status:          InvoiceStatusUnpaid,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		},
		{
			name: "case2: normal with margin error",
			in: in{
				companyID:       1,
				companyClientID: 1,
				issueAt:         mtime.Now(),
				paymentAmount:   10010,
				dueAt:           mtime.Now().AddDate(0, 1, 0),
			},
			want: &Invoice{
				ID:              0,
				RandID:          "inv-id",
				CompanyID:       1,
				CompanyClientID: 1,
				IssueAt:         now,
				PaymentAmount:   10010,
				BillingAmount:   10450, // 10450.44
				Fee:             400,
				FeeRatio:        0.04,
				Tax:             40,
				TaxRatio:        0.1,
				DueAt:           now.AddDate(0, 1, 0),
				Status:          InvoiceStatusUnpaid,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		},
		{
			name: "case3: validation error",
			in: in{
				companyID:       1,
				companyClientID: 1,
				issueAt:         mtime.Now(),
				paymentAmount:   1000,
				dueAt:           time.Time{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewInvoiceBuilder().
				CompanyID(tt.in.companyID).
				CompanyClientID(tt.in.companyClientID).
				IssueAt(tt.in.issueAt).
				PaymentAmount(tt.in.paymentAmount).
				DueAt(tt.in.dueAt).
				Build()

			if (err != nil) != tt.wantErr {
				t.Errorf("InvoiceBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}
}
