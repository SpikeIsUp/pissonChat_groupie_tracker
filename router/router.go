package router

import (
	"net/http"

	"github.com/SpikeIsUp/igApi/controller"
)

func Router() *http.ServeMux {
	mux := http.Router()

	// routes
	mux.HandlFunc("/aPropos", controller.APropos)
	mux.HandlFunc("/categories", controller.Categories)
	mux.HandlFunc("/collection", controller.Collection)
	mux.HandlFunc("/favoris", controller.Favoris)
	mux.HandlFunc("/", controller.Home)
	mux.HandlFunc("/recherche", controller.Recherche)
	mux.HandlFunc("/ressources", controller.Ressources)
}