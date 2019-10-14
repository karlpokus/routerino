package routest

import (
	"testing"
	"net/http"
)

func hi() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	}
}

func TestRoutes(t *testing.T) {
  Test(t, []Data{
    {
      "hi",
      nil,
      hi(),
      200,
      []byte("hi"),
    },
  })
}
