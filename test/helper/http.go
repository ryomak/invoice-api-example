package helper

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/ryomak/invoice-api-example/infrastructure/env"
	"github.com/ryomak/invoice-api-example/presentation/router"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestServer struct {
	*httptest.Server
	Shutdown func() error
}

func RunTestServer() (*TestServer, error) {
	// test用のenvを設定
	if err := env.Build(); err != nil {
		return nil, err
	}

	manager, err := NewDBContainerManager()
	if err != nil {
		return nil, err
	}
	conn, shutdown, err := manager.CreateOrGetConn()
	if err != nil {
		return nil, err
	}
	if err := TestDBSetup(conn); err != nil {
		return nil, err
	}

	r, err := router.New(conn)
	if err != nil {
		return nil, err
	}
	r.Routes()

	return &TestServer{
		Server:   httptest.NewServer(r),
		Shutdown: shutdown,
	}, nil
}

func (ts *TestServer) TryRequest(t *testing.T, desc, method, path string, headers map[string]string, payload string, wantCode int, wantBody string) {

	req, err := http.NewRequest(method, ts.URL+path, strings.NewReader(payload))
	if err != nil {
		t.Errorf("%s: generate request: %v", desc, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	c := http.DefaultClient

	resp, err := c.Do(req)
	if err != nil {
		t.Errorf("%s: http.Get: %v", desc, err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%s: reading body: %v", desc, err)
		return
	}

	if resp.StatusCode != wantCode {
		t.Errorf("%s: got HTTP %d, want %d", desc, resp.StatusCode, wantCode)
		t.Errorf("response body: %s", string(body))
		return
	}

	if wantBody != "" {
		if eq, o1, o2, err := isEqualJSON(string(body), wantBody); err != nil {
			t.Errorf("unmershalJSON %s", err.Error())
		} else if !eq {
			t.Errorf("%s: body is different", desc)
			t.Log(cmp.Diff(o1, o2))
			t.Log(string(body))
		}
		return
	}
	t.Log(desc, ":ok")
}

func isEqualJSON(s1, s2 string) (bool, any, any, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, nil, nil, fmt.Errorf("Error mashalling string 1 :: %w ", err)
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, nil, nil, fmt.Errorf("Error mashalling string 2 :: %w ", err)
	}
	return cmp.Equal(o1, o2), o1, o2, nil
}
