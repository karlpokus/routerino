package routest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/karlpokus/srv"
)

func hi(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi %s", s)
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Fprintf(w, "hello %s", params.ByName("user"))
}

func TestRoute(t *testing.T) {
	Test(t, nil, []Data{
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

func TestRouter(t *testing.T) {
	Test(t, func() http.Handler {
		router := httprouter.New()
		router.HandlerFunc("GET", "/greet/:user", Greet)
		return router
	}, []Data{
		{
			"Greet from router",
			"GET",
			"/greet/bob",
			nil,
			nil,
			200,
			[]byte("hello bob"),
		},
	})
}

func TestServer(t *testing.T) {
	Test(t, func() http.Handler {
		s, _ := srv.New(func(s *srv.Server) error {
			router := s.DefaultRouter()
			router.Handle("/hi", hi("bob"))
			s.Router = router
			s.Quiet()
			return nil
		})
		return s
	}, []Data{
		{
			"hi from server",
			"GET",
			"/hi",
			nil,
			nil,
			200,
			[]byte("hi bob"),
		},
	})
}

func ExampleTest_route(t *testing.T) {
	Test(t, nil, []Data{
		{
			"hi",
			"GET",
			"/",
			nil,
			hi("bob"),
			200,
			[]byte("hi bob"),
		},
	})
}

func ExampleTest_router(t *testing.T) {
	Test(t, func() http.Handler {
		router := httprouter.New()
		router.HandlerFunc("GET", "/greet/:user", Greet)
		return router
	}, []Data{
		{
			"Greet",
			"GET",
			"/greet/bob",
			nil,
			nil, // use router as handler
			200,
			[]byte("hello bob"),
		},
	})
}
