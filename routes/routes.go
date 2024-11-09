package routes

import (
	"github.com/gorilla/mux"
	"pismo/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Account routes
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", controllers.GetAccount).Methods("GET")

	// Transaction routes
	router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	return router
}
