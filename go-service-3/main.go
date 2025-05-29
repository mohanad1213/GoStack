package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Service 3 🚀!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Service running on port 8000")
	http.ListenAndServe(":8000", nil)
}
