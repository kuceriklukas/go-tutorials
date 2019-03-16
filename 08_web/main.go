package main

import (
	"fmt"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "<h1>This is awesome!</h1>")
}

func about(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "<h1>This is still awesome!</h1><h2>But it is about the awesome things</h2>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting...")

	http.ListenAndServe(":3003", nil)
}
