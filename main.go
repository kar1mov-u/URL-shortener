package main

import (
	"net/http"

	"github.com/kar1mov-u/URL-shortener/db"
	"github.com/kar1mov-u/URL-shortener/src"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", src.Ind)
	mux.HandleFunc("/shorten", src.Shorten)
	mux.HandleFunc("/long/", src.Redirect)
	http.ListenAndServe(":8000", mux)

}
