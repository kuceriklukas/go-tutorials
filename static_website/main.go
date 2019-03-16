package main

import (
	"log"
	"net/http"
)

const host = "http://localhost"
const port = ":3003"

func main() {
	log.Println("Server starting...")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening on " + host + port)

	http.ListenAndServe(port, nil)
}
