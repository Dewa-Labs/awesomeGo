package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Gopher! Welcome to your server.")
	})

	// Optional: a health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	// Start the server on 0.0.0.0:8080
	addr := "0.0.0.0:8080"
	fmt.Println("Server listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
