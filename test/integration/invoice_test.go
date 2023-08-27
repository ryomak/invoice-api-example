package integration

import (
	"fmt"
	mtime "github.com/ryomak/invoice-api-example/pkg/time"
	"github.com/ryomak/invoice-api-example/pkg/unique"
	"github.com/ryomak/invoice-api-example/test/helper"
	"net/http"
	"testing"
	"time"
)

func TestInvoice(t *testing.T) {
	t.Parallel()
	ts, err := helper.RunTestServer()
	if err != nil {
		t.Errorf("runTestServer: %v", err)
		return
	}
	defer ts.Close()
	defer ts.Shutdown()

	t.Run("create invoice", func(t *testing.T) {
		id := "testidxxx"
		unique.SetFakeGenerateID(id)
		now := time.Date(2023, 8, 1, 15, 0, 0, 0, time.UTC)
		mtime.SetFakeNow(now)

		headers := map[string]string{
			"X-Debug-Id": "test_user",
		}
		payload := `{
    "companyClientRandId": "test_client",
    "paymentAmount": 10000,
    "dueAt": "2023-09-01T15:00:00Z"
}`
		ts.TryRequest(t, "create invoice", http.MethodPost, "/api/invoices", headers, payload, 200, fmt.Sprintf(`{
   "invoice":{
      "randId":"inv-%s",
      "status":"unpaid",
      "issueAt":"2023-08-01T15:00:00Z",
      "paymentAmount":10000,
      "billingAmount":10440,
      "fee":400,
      "feeRatio":0.04,
      "tax":40,
      "taxRatio":0.1,
      "dueAt":"2023-09-01T15:00:00Z"
   }
}`, id))
	})

	t.Run("get invoice", func(t *testing.T) {
		now := time.Date(2023, 8, 1, 15, 0, 0, 0, time.UTC)
		mtime.SetFakeNow(now)

		url := "/api/invoices?from=2022-09-01T15:00:00Z&to=2022-11-01T15:00:00Z&limit=3&offset=0"

		headers := map[string]string{
			"X-Debug-Id": "test_user",
		}
		ts.TryRequest(t, "get invoice", http.MethodGet, url, headers, "", 200, `{
   "items":[
      {
         "randId":"inv-cjlk6r5315okaj305qd0",
         "status":"paid",
         "issueAt":"2022-10-01T15:00:00Z",
         "paymentAmount":10000,
         "billingAmount":10440,
         "fee":400,
         "feeRatio":0.04,
         "tax":40,
         "taxRatio":0.1,
         "dueAt":"2022-10-01T15:00:00Z"
      }
   ]
}`)
	})
}
