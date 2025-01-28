package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", src.ind)
	mux.HandleFunc("/shorten", src.shorten)
	mux.HandleFunc("/long/", src.redirect)
	http.ListenAndServe(":8000", mux)

}
