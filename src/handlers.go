package src

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
var hmap = make(map[string]string)

func ind(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Cannot load template", http.StatusInternalServerError)
	}

}

func shorten(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("url")
	if longUrl == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
	}

	shortUrl := hashing(longUrl)
	hmap[shortUrl] = longUrl
	data := map[string]string{"ShortURL": shortUrl}
	if err := templates.ExecuteTemplate(w, "result.html", data); err != nil {
		http.Error(w, "Cannot load template", http.StatusInternalServerError)
	}

}

func hashing(url string) string {
	hash := sha256.Sum256([]byte(url))
	short := base64.URLEncoding.EncodeToString(hash[:])
	return short[:8]
}

func redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/long/"):]
	fmt.Println(path)
	longUrl := hmap[path]
	fmt.Println(longUrl)

	fmt.Printf("Redirecting short URL %s to long URL %s\n", path, longUrl)
	http.Redirect(w, r, longUrl, http.StatusFound)
}
