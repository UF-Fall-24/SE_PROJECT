package main

import (
	"book-ease-backend/config"
	"book-ease-backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Connect to the database
	config.Connect()

	// Use the router for routes
	router := routes.SetupRoutes()

	// Add CORS support
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the server with CORS middleware
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", cors(router)))
}
