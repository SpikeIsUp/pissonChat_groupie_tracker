package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "favorites.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS favorites (
		id TEXT PRIMARY KEY,
		name TEXT,
		url TEXT
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
