# routest
Fast and easy testing of your http api. Works with the stdlibs `testing` pkg.

[![GoDoc](https://godoc.org/github.com/karlpokus/routest?status.svg)](https://godoc.org/github.com/karlpokus/routest)

# usage
```go
import (
	"testing"
	"net/http"
	"github.com/karlpokus/routest"
)

func hi() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	}
}

func TestRoutes(t *testing.T) {
  routest.Test(t, []routest.Data{
    {
      "hi",
      nil,
      hi(),
      200,
      []byte("hi"),
    },
  })
}
```

# license
MIT
