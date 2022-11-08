package main

import (
	"github.com/gorilla/mux"
	"github.com/joeernest/gormapirest/handlers"
	"github.com/joeernest/gormapirest/models"
	"log"
	"net/http"
)

func main() {

	models.MigrateUsers()
	router := mux.NewRouter()

	// Routing
	router.HandleFunc("/api/user/", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// Starting server
	log.Fatal(http.ListenAndServe(":3000", router))
}
