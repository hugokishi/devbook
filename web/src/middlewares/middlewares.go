package middlewares

import (
	"log"
	"net/http"
	"web/src/cookies"
)

// Logger - Log the requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authentication - Verify authentication data
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		next(w, r)
	}
}
