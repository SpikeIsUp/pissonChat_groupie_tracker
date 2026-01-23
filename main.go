package main

import (
	"log"
	"net/http"

	"github.com/SpikeIsUp/pissonChat_groupie_tracker/SQLiteinternal/storage"
	"github.com/SpikeIsUp/pissonChat_groupie_tracker/router"
)

func main() {
	storage.InitDB()

	mux := router.SetupRouter()

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
