package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func Requete(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Requete")
}

func Individuel(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Individuel")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./Templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Requete)
	http.HandleFunc("/Individuel", Individuel)

	fmt.Println("http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
