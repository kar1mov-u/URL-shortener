package src

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kar1mov-u/URL-shortener/db"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Ind(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Cannot load template", http.StatusInternalServerError)
	}

}

func Shorten(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("url")
	if longUrl == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
	}
	//CHeck if the longUrl in database
	var shortUrl string

	getLongQuery := `SELECT short_url FROM main WHERE long_url = ?`
	err := db.DB.QueryRow(getLongQuery, longUrl).Scan(&shortUrl)
	//If no record, create one
	if err == sql.ErrNoRows {
		shortUrl = Hashing(longUrl)
		insertQuery := `INSERT INTO main (long_url, short_url, created_at) VALUES (?,?,CURRENT_TIMESTAMP);`
		if _, err := db.DB.Exec(insertQuery, longUrl, shortUrl); err != nil {
			panic(err)
		}

	} else if err != nil {
		fmt.Printf("failed to fetch data %v", err)
	}

	data := map[string]string{"ShortURL": shortUrl}
	if err := templates.ExecuteTemplate(w, "result.html", data); err != nil {
		http.Error(w, "Cannot load template", http.StatusInternalServerError)
	}

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	var long_url string
	shortUrl := r.URL.Path[len("/long/"):]
	getLongQuery := "SELECT long_url FROM main WHERE short_url = ?;"
	err := db.DB.QueryRow(getLongQuery, shortUrl).Scan(&long_url)
	if err != nil {
		http.Error(w, "Bad URL", http.StatusNotFound)
		return
	}

	fmt.Printf("Redirecting short URL %s to long URL %s\n", shortUrl, long_url)
	http.Redirect(w, r, long_url, http.StatusFound)
}
