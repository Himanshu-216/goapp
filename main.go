package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "Kubernetes Cluster AWS"
	}
	fmt.Fprintf(w, "Hello and Welcome from %s!", serviceName)
}

func main() {
	mux := http.NewServeMux()

	// Root route ONLY
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	})

	port := ":8080"
	fmt.Println("Server running at", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
