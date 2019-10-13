package routerino

import (
  "bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Table struct {
  Name string
	RequestBody io.Reader
	Handler http.HandlerFunc
	Status int
	ResponseBody []byte
}

func Test(t *testing.T, table []Table) {
  for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", tt.RequestBody)
			w := httptest.NewRecorder()
			tt.Handler(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)
			if res.StatusCode != tt.Status {
				t.Errorf("expected %d, got %d", tt.Status, res.StatusCode)
			}
			if !bytes.Equal(bytes.TrimSpace(body), tt.ResponseBody) {
				t.Errorf("expected %s, got %s", tt.ResponseBody, bytes.TrimSpace(body))
			}
		})
	}
}
