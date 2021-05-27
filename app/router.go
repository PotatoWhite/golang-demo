package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUsers).Methods("POST")
	//
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")

	log.Println("Application Initialized")
	log.Fatal(http.ListenAndServe(":8080", r))
}

