package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to connect database %v", err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS main(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	long_url TEXT NOT NULL,
	short_url TEXT NOT NULL,
	created_at DATETIME
	);
	`
	if _, err = DB.Exec(createTable); err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
}
