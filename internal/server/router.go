package server

import (
	"fmt"
	"net/http"
)

// NewRouter builds and returns an http.Handler
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	// Root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Gopher! Welcome to your server.")
	})

	return mux
}
