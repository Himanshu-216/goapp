package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	serviceName := os.Getenv("SERVICE_NAME") // read from env variable
	if serviceName == "" {
		serviceName = "Kubernetes Cluster aws"
	}
	fmt.Fprintf(w, "Hello and Welcome from %s!", serviceName)
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	fmt.Println("Server running at", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
