package main

import (
	"log"
	"net/http"

	"pismo/account"
	"pismo/controllers"
	"pismo/database"
	"pismo/routes"
	"pismo/transaction"
)

func main() {
	// Initialize the database
	db, errDB := database.ConnectDatabase("main.db")
	if errDB != nil {
		log.Fatal(errDB)
	}

	accountSvc := account.NewAccountService(db)
	txnSvc := transaction.NewService(db, accountSvc)

	txnController := controllers.NewTransactionController(txnSvc)
	accountController := controllers.NewAccountController(accountSvc)

	// Set up the router
	router := routes.SetupRouter(txnController, accountController)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
