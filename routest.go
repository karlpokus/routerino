// Package routest provides a quick and easy way of testing your http api
package routest

import (
  "bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Data struct {
  Name string
	RequestBody io.Reader
	Handler http.HandlerFunc
	Status int
	ResponseBody []byte
}

// Test runs assertion tests on data and reports to t
func Test(t *testing.T, data []Data) {
  for _, d := range data {
		t.Run(d.Name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", d.RequestBody)
			w := httptest.NewRecorder()
			d.Handler(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)
			if res.StatusCode != d.Status {
				t.Errorf("expected %d, got %d", d.Status, res.StatusCode)
			}
			if !bytes.Equal(bytes.TrimSpace(body), d.ResponseBody) {
				t.Errorf("expected %s, got %s", d.ResponseBody, bytes.TrimSpace(body))
			}
		})
	}
}
