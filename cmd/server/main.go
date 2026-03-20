package main

import (
	"fmt"
	"net/http"
	"os"

	"awesomeGo/internal/server"
)

func main() {
	router := server.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port
	fmt.Println("Server listening on", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		fmt.Println("Server error:", err)
	}
}
