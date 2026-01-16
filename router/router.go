package router

import (
	"net/http"

	"pissonChat_groupie_tracker/controller"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/search", controller.Search)
	mux.HandleFunc("/favorites", controller.Favorites)
	mux.HandleFunc("/favorite/add", controller.AddFavorite)
	mux.HandleFunc("/favorite/remove", controller.RemoveFavorite)
	mux.HandleFunc("/about", controller.About)

	return mux
}
