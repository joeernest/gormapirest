package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joeernest/gormapirest/models"
	"github.com/joeernest/gormapirest/routing"
	"log"
	"net/http"
)

func main() {

	models.MigrateUsers()

	router := mux.NewRouter().StrictSlash(true)
	routing.RegisterUserRoutes(router)

	// Starting server
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port 3000"))
	log.Fatal(http.ListenAndServe(":3000", router))
}
