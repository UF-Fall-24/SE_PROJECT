package main

import (
	"book-ease-backend/config"
	"book-ease-backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv" // Import dotenv package
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using defaults")
	}

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
	fmt.Println("Server running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", cors(router)))
}
