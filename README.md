# routest
Fast and easy way of testing your http api. Works with the stdlibs `testing` pkg.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/karlpokus/routest/v2@v2.1.0)

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
		w.Header().Set("Etag", "abc")
		fmt.Fprintf(w, "hi %s", s)
	}
}

func TestRoute(t *testing.T) {
	routest.Test(t, nil, []routest.Data{
		{
			Name: "hi from route",
			Path: "/",
			Handler: hi("bob"),
			Status: 200,
			ResponseBody: []byte("hi bob"),
			ResponseHeader: http.Header{
				"Etag": []string{"abc"},
			},
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
			Name: "Greet from router",
			Method: "GET",
			Path: "/greet/bob",
			Status: 200,
			ResponseBody: []byte("hello bob"),
		},
	})
}
```

# todos
- [x] allow for custom server/router to register as handler
- [x] fix v2 module path
- [x] Add http.Header to Data
- [x] Make some Data fields optional

# license
MIT
