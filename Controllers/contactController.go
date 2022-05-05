package controllers

import (
	a "assignment/Config"
	model "assignment/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var con [](model.Contact)
	a.DB.First(&con, params["id"])
	json.NewEncoder(w).Encode(con)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var con [](model.Contact)
	a.DB.First(&con, params["id"])
	json.NewDecoder(r.Body).Decode(&con)
	a.DB.Save(&con)
	json.NewEncoder(w).Encode(con)
}
