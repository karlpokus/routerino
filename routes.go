package routerino

import "net/http"

func hi() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	}
}
