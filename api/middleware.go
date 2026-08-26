package api

import (
	"log"
	"net/http"
	"time"
)

func withRequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
func jsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
func method(allowed ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, m := range allowed {
				if r.Method == m {
					next.ServeHTTP(w, r)
					return
				}
			}
			w.Header().Set("Allow", allowed[0])
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		})
	}
}
