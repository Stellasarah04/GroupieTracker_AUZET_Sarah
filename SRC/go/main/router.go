package main

import (
	"github.com/gorilla/mux"
	"https://github.com/Stellasarah04/GroupieTracker_AUZET_Sarah/tree/main/SRC/go/controler"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /pays/ to /pays
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/pays").Name("Index").HandlerFunc(controllers.PaysIndex)
	router.Methods("POST").Path("/pays").Name("Create").HandlerFunc(controllers.PaysCreate)
	router.Methods("GET").Path("/pays/{id}").Name("Show").HandlerFunc(controllers.PaysShow)
	router.Methods("PUT").Path("/pays/{id}").Name("Update").HandlerFunc(controllers.PaysUpdate)
	router.Methods("DELETE").Path("/pays/{id}").Name("DELETE").HandlerFunc(controllers.PaysDelete)
	return router
}
