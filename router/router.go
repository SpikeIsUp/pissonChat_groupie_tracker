package router

import (
	"net/http"

	"github.com/SpikeIsUp/pissonChat_groupie_tracker/controller"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Accueil
	mux.HandleFunc("/", controller.Home)

	// Recherche
	mux.HandleFunc("/search", controller.Search)

	// Favoris
	mux.HandleFunc("/favorites", controller.Favorites)

	// Ajouter un favori
	mux.HandleFunc("/favorite/add", controller.AddFavorite)

	// À propos
	mux.HandleFunc("/about", controller.About)

	// Static files (CSS, JS, images)
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
