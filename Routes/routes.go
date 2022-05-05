package routes

import (
	a "assignment/Controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter() // created a router

	// implementing different routes
	r.HandleFunc("/users", a.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", a.GetUser).Methods("GET")
	r.HandleFunc("/users", a.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", a.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", a.DeleteUser).Methods("DELETE")

	r.HandleFunc("/address/{id}", a.GetAddress).Methods("GET")
	r.HandleFunc("/address/{id}", a.UpdateAddress).Methods("PUT")

	r.HandleFunc("/contact/{id}", a.GetContact).Methods("GET")
	r.HandleFunc("/contact/{id}", a.UpdateContact).Methods("PUT")

	fmt.Println("Routes and Controllers have been established")
	fmt.Println("Use Postman for checking")

	// Have to link the tables, foreign key

	log.Fatal(http.ListenAndServe(":9000", r)) // (port, router)
}
