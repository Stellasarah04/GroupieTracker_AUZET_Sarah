package main

import (
	"log"
	"net/http"

	"github.com/synbioz/go_api/config"
	"github.com/synbioz/go_api/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewCar(&models.Pays{Manufacturer: "citroen", Design: "ds3", Style: "sport", Doors: 4})

	log.Fatal(http.ListenAndServe(":8080", router))
}
