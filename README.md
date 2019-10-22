# routest
Fast and easy way of testing your http api. Works with the stdlibs `testing` pkg.

[![GoDoc](https://godoc.org/github.com/karlpokus/routest?status.svg)](https://godoc.org/github.com/karlpokus/routest)

# install
```bash
$ go get github.com/karlpokus/routest/v2
```

# usage
Test a route
```go
import (
	// ...
	"github.com/karlpokus/routest/v2"
)

func hi(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi %s", s)
	}
}

func TestRoute(t *testing.T) {
	routest.Test(t, nil, []routest.Data{
		{
			"hi from route",
			"GET",
			"/",
			nil,
			hi("bob"),
			200,
			[]byte("hi bob"),
		},
	})
}
```
Test registered routes
```go
import (
	// ...
	"github.com/karlpokus/routest/v2"
	"github.com/julienschmidt/httprouter"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Fprintf(w, "hello %s", params.ByName("user"))
}

func TestRouter(t *testing.T) {
	routest.Test(t, func() http.Handler {
		router := httprouter.New()
		router.HandlerFunc("GET", "/greet/:user", Greet)
		return router
	}, []routest.Data{
		{
			"Greet from router",
			"GET",
			"/greet/bob",
			nil,
			nil, // use router as handler
			200,
			[]byte("hello bob"),
		},
	})
}
```

# todos
- [x] allow for custom server/router to register as handler
- [x] fix v2 module path

# license
MIT
