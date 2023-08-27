package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/ryomak/invoice-api-example/pkg/logger"
	"io"
	"net/http"
)

func AccessLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// body出力
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		bodies := map[string]interface{}{}
		_ = json.Unmarshal(body, &bodies)

		// パスワードはマスクする
		if bodies["password"] != nil {
			bodies["password"] = "xxx"
		}

		next.ServeHTTP(w, r)
		logger.InfofWithData(r.Context(), bodies, "[%s]%s", r.Method, r.URL.Path)
	}
	return http.HandlerFunc(fn)
}
