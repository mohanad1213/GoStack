package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Service 1 🚀!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Service running on port 8000")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
