package controller

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/SpikeIsUp/pissonChat_groupie_tracker/ApiMemeMakerinternal/meme"
	"github.com/SpikeIsUp/pissonChat_groupie_tracker/SQLiteinternal/storage"
)

func Home(w http.ResponseWriter, r *http.Request) {
	memes, err := meme.GetMemes()
	if err != nil {
		http.Error(w, "Erreur API", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("template/home.html"))
	if err := tmpl.Execute(w, memes); err != nil {
		log.Println("Erreur template Home :", err)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	memes, err := meme.GetMemes()
	if err != nil {
		http.Error(w, "Erreur API", http.StatusInternalServerError)
		return
	}

	var result []meme.Meme
	for _, m := range memes {
		if strings.Contains(strings.ToLower(m.Name), strings.ToLower(query)) {
			result = append(result, m)
		}
	}

	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println("Erreur template Search :", err)
	}
}

func Favorites(w http.ResponseWriter, r *http.Request) {
	memes, err := storage.GetFavorites()
	if err != nil {
		http.Error(w, "Erreur DB", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	if err := tmpl.Execute(w, memes); err != nil {
		log.Println("Erreur template Favorites :", err)
	}
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("name")
	blank := r.FormValue("blank")

	if err := storage.AddFavorite(id, name, blank); err != nil {
		http.Error(w, "Impossible d'ajouter aux favoris", http.StatusInternalServerError)
		return
	}

	log.Printf("Favori ajouté : %s (%s)", name, id)
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	id := r.FormValue("id")

	if err := storage.RemoveFavorite(id); err != nil {
		http.Error(w, "Impossible de supprimer le favori", http.StatusInternalServerError)
		return
	}

	log.Printf("Favori supprimé : %s", id)
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/aPropos.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Erreur template About :", err)
	}
}
