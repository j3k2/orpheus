package main

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/capture", Capture)

	mux.HandleFunc("GET /files/", ListFiles)
	mux.HandleFunc("GET /files/{filename}", GetFile)

	gui := http.FileServer(http.Dir("./gui"))
	mux.Handle("/", gui)

	return loggingMiddleware(mux)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}