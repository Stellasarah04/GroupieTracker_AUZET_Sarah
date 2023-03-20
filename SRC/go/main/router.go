package main

import (
	"github.com/gorilla/mux"
	"github.com/synbioz/go_api/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/pays").Name("Index").HandlerFunc(controllers.CarsIndex)
	router.Methods("POST").Path("/pays").Name("Create").HandlerFunc(controllers.CarsCreate)
	router.Methods("GET").Path("/pays/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
	router.Methods("PUT").Path("/pays/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
	router.Methods("DELETE").Path("/pays/{id}").Name("DELETE").HandlerFunc(controllers.CarsDelete)
	return router
}
