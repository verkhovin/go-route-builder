package routebuilder

import "net/http"

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404 Not Found"))
	})
}
