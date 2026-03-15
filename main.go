package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}