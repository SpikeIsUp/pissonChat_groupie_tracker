package main

import (
	"log"
	"net/http"

	"pissonChat_groupie_tracker/internal/storage"
	"pissonChat_groupie_tracker/router"
)

func main() {
	storage.InitDB()

	r := router.SetupRouter()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
