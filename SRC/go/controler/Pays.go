package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"https://github.com/Stellasarah04/GroupieTracker_AUZET_Sarah/tree/main/SRC/go/models"
)

func PaysIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllPays())
}

func PaysCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var pays models.Pays

	err = json.Unmarshal(body, &pays)

	if err != nil {
		log.Fatal(err)
	}

	models.NewPays(&pays)

	json.NewEncoder(w).Encode(pays)
}

func PaysShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	pays := models.FindPaysById(id)

	json.NewEncoder(w).Encode(pays)
}

func PaysUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	pays := models.FindPaysById(id)

	err = json.Unmarshal(body, &pays)

	models.UpdatePays(pays)

	json.NewEncoder(w).Encode(pays)
}
func PaysDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = models.DeletePaysById(id)
}
