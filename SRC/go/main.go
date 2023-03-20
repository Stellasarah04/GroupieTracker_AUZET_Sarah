package main

import (
	"encoding/json"
	"fmt"
)

type Pays struct {
	Name      string `json : "name"`
	Capital   string `json : "Capital"`
	Languages string `json : "Languages"`
	Monnaie   string `json : "Monnaie"`
	Continent string `json : "Continent"`
}

func main() {
	jsonFromAPI := `
[
	{
		"Name" : "Japon",
		"Capital" : "Tokyo",
		"Languages" : "Japonais", 
		"Monnaie" : "Yen", 
		"Continent" : "Asie"
	}
]`

	var pays []Pays

	err := json.Unmarshal([]byte(jsonFromAPI), &pays)

	if err != nil {
		fmt.Println("Error Unmarshalling json", err)
	}

	fmt.Println("json : %v\n", pays)
}

/*import (
	"loUnmarshal
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
}*/
