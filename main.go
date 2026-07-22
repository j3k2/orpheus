package main


import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Port for HTTP server")
	flag.Parse()

	r := NewRouter()

	addr := fmt.Sprintf(":%s", *port)

	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}