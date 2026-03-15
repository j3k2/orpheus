package main

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/capture", CaptureHandler)
    fs := http.FileServer(http.Dir("./static"))
    mux.Handle("/", fs)
	return loggingMiddleware(mux)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}