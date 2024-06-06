package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s - path: %s", r.Method, r.URL.EscapedPath())
		next(w, r)
	}
}
