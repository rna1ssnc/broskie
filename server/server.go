package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
