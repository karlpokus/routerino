// Package routest provides a quick and easy way of testing your http api
package routest

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Only Name, Method and Path are required. All other fields are optional.
// A non-nil RegisterFunc overrides Handler.
type Data struct {
	Name, Method, Path string
	RequestBody        io.Reader
	RequestHeader      http.Header
	Handler            http.Handler
	Status             int
	ResponseBody       []byte
	ResponseHeader     http.Header
}

type RegisterFunc func() http.Handler

// Test runs assertion tests on data and reports to t
func Test(t *testing.T, fn RegisterFunc, data []Data) {
	var handler http.Handler
	if fn != nil {
		handler = fn()
	}
	for _, d := range data {
		t.Run(d.Name, func(t *testing.T) {
			r := httptest.NewRequest(d.Method, d.Path, d.RequestBody)
			if d.RequestHeader != nil {
				for k, v := range d.RequestHeader {
					r.Header.Set(k, join(v))
				}
			}
			w := httptest.NewRecorder()
			if handler == nil {
				handler = d.Handler
			}
			handler.ServeHTTP(w, r)
			res := w.Result()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Read body err: %s", err)
			}
			if d.Status != 0 {
				if res.StatusCode != d.Status {
					t.Errorf("expected %d, got %d", d.Status, res.StatusCode)
				}
			}
			if len(d.ResponseBody) != 0 {
				if !bytes.Equal(bytes.TrimSpace(body), d.ResponseBody) {
					t.Errorf("expected %s, got %s", d.ResponseBody, bytes.TrimSpace(body))
				}
			}
			if d.ResponseHeader != nil {
				for k, v1 := range d.ResponseHeader {
					v2, ok := res.Header[k]
					if !ok {
						t.Errorf("%s is missing from response header", k)
					}
					s1, s2 := join(v1), join(v2)
					if s1 != s2 {
						t.Errorf("%s is not equal to %s for response header %s", s1, s2, k)
					}
				}
			}
		})
	}
}

func join(list []string) string {
	return strings.Join(list, ",")
}
