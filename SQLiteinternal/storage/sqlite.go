package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initialise la base de données et crée la table favorites si elle n'existe pas
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal("Erreur d'ouverture de la DB :", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS favorites (
		id TEXT PRIMARY KEY,
		name TEXT,
		blank TEXT
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Erreur création table favorites :", err)
	}
}

// AddFavorite ajoute un meme aux favoris (ignore les doublons)
func AddFavorite(id, name, blank string) error {
	if id == "" || name == "" || blank == "" {
		return nil
	}
	_, err := DB.Exec("INSERT OR IGNORE INTO favorites (id, name, blank) VALUES (?, ?, ?)", id, name, blank)
	if err != nil {
		log.Println("Erreur ajout favori :", err)
	}
	return err
}

// RemoveFavorite supprime un meme des favoris
func RemoveFavorite(id string) error {
	_, err := DB.Exec("DELETE FROM favorites WHERE id = ?", id)
	if err != nil {
		log.Println("Erreur suppression favori :", err)
	}
	return err
}

// GetFavorites récupère tous les favoris
func GetFavorites() ([]Meme, error) {
	rows, err := DB.Query("SELECT id, name, blank FROM favorites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memes []Meme
	for rows.Next() {
		var m Meme
		if err := rows.Scan(&m.ID, &m.Name, &m.Blank); err != nil {
			return nil, err
		}
		memes = append(memes, m)
	}
	return memes, nil
}

// Meme struct utilisé pour les favoris
type Meme struct {
	ID    string
	Name  string
	Blank string
}
