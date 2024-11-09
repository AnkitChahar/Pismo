package main

import (
	"log"
	"net/http"

	"pismo/database"
	"pismo/routes"
)

func main() {
	// Initialize the database
	database.ConnectDatabase()

	// Set up the router
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
