package routing

import (
	"github.com/gorilla/mux"
	"github.com/joeernest/gormapirest/handlers"
)

func RegisterUserRoutes(router *mux.Router) {
	// User routing
	router.HandleFunc("/api/user/", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUserByID).Methods("DELETE")

}
