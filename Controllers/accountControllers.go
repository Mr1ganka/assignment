package controllers

import (
	a "assignment/Config"
	model "assignment/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// handler methods
// http.ResponseWriter -> sends data to the HTTP client
// http.Request -> data structure that represent a client HTTP request
// *http.Request -> pointer of the HTTP request
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users [](model.Account)
	a.DB.Preload("Address").Preload("Contact").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](model.Account)
	a.DB.Preload("Address").Preload("Contact").First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user [](model.Account)
	json.NewDecoder(r.Body).Decode(&user)
	a.DB.Preload("Address").Preload("Contact").Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](model.Account)
	a.DB.Preload("Address").Preload("Contact").First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	a.DB.Preload("Address").Preload("Contact").Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](model.Account)
	a.DB.Preload("Address").Preload("Contact").Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
