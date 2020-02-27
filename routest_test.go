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
		if r.Header.Get("serious") == "error" {
			http.Error(w, "error", 500)
			return
		}
		w.Header().Set("Etag", "abc")
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
			Name:         "hi from route",
			Method:       "GET",
			Path:         "/",
			Handler:      hi("bob"),
			Status:       200,
			ResponseBody: []byte("hi bob"),
			ResponseHeader: http.Header{
				"Etag": []string{"abc"},
			},
		},
		{
			Name:   "Trigger error",
			Method: "GET",
			Path:   "/",
			RequestHeader: http.Header{
				"serious": []string{"error"},
			},
			Handler:      hi(""),
			Status:       500,
			ResponseBody: []byte("error"),
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
			Name:         "Greet from router",
			Method:       "GET",
			Path:         "/greet/bob",
			ResponseBody: []byte("hello bob"),
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
			nil,
			200,
			[]byte("hi bob"),
			nil,
		},
	})
}

func ExampleTest_route(t *testing.T) {
	Test(t, nil, []Data{
		{
			Name:         "hi",
			Method:       "GET",
			Path:         "/",
			Handler:      hi("bob"),
			Status:       200,
			ResponseBody: []byte("hi bob"),
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
			Name:         "Greet",
			Method:       "GET",
			Path:         "/greet/bob",
			Status:       200,
			ResponseBody: []byte("hello bob"),
		},
	})
}
