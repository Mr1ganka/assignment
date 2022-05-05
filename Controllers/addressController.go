package controllers

import (
	a "assignment/Config"
	model "assignment/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var address [](model.Address)
	a.DB.Find(&address)
	json.NewEncoder(w).Encode(address)
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var address [](model.Address)
	a.DB.First(&address, params["id"])
	json.NewDecoder(r.Body).Decode(&address)
	a.DB.Save(&address)
	json.NewEncoder(w).Encode(address)
}
