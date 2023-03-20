package main

import (
	"log"
	"net/http"

	"https://github.com/Stellasarah04/GroupieTracker_AUZET_Sarah/tree/main/SRC/go/config"
	"https://github.com/Stellasarah04/GroupieTracker_AUZET_Sarah/tree/main/SRC/go/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewPays(&models.Pays{Name: "Japon", Capital: "Tokyo", Monnaie: "Yen", Languages: "Japonais"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
