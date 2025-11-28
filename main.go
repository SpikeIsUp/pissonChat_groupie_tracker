package main

import(
	"net/http"
	"log"
)

func main() {

	// r devient router
	r := router.Router

// fichiers statiques
	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.Handle("/", r)

	log.Println("Serveur lancé--> http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
