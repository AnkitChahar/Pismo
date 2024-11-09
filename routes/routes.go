package routes

import (
	"github.com/gorilla/mux"
	"pismo/controllers"
)

func SetupRouter(txnController *controllers.TransactionController, accountController *controllers.AccountController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Account routes
	router.HandleFunc("/accounts", accountController.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountController.GetAccount).Methods("GET")

	// Transaction routes
	router.HandleFunc("/transactions", txnController.CreateTransaction).Methods("POST")

	return router
}
