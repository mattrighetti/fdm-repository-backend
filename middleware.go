package main

import (
	"log"
	"net/http"
)

// loggerMiddleware logs requests against the server
func loggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v -> %v: %v", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
