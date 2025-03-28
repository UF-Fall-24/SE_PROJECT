package routes

import (
	"book-ease-backend/controllers"
	"book-ease-backend/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Public package endpoints: get all packages and get a package by id
	router.HandleFunc("/packages", controllers.GetPackages).Methods("GET")
	router.HandleFunc("/packages/{id}", controllers.GetPackage).Methods("GET")

	// endpoint: fetch hotels by location using query parameter 'location'
	router.HandleFunc("/hotels/location", controllers.GetHotelsByLocation).Methods("GET")
	// Public hotel endpoints
	router.HandleFunc("/hotels", controllers.GetHotels).Methods("GET")
	router.HandleFunc("/hotels/{id}", controllers.GetHotel).Methods("GET")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTAuth)

	protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")

	// User Profile Management Routes
	protected.HandleFunc("/users/profile/{id}", controllers.GetUserProfile).Methods("GET")
	protected.HandleFunc("/users/profile/{id}", controllers.UpdateUserProfile).Methods("PUT")

	// Package endpoints
	protected.HandleFunc("/packages", controllers.CreatePackage).Methods("POST")
	protected.HandleFunc("/packages/{id}", controllers.UpdatePackage).Methods("PUT")
	protected.HandleFunc("/packages/{id}", controllers.DeletePackage).Methods("DELETE")

	// Protected hotel endpoints
	protected.HandleFunc("/hotels", controllers.CreateHotel).Methods("POST")
	protected.HandleFunc("/hotels/{id}", controllers.UpdateHotel).Methods("PUT")
	protected.HandleFunc("/hotels/{id}", controllers.DeleteHotel).Methods("DELETE")

	// Accommodation Endpoints
	protected.HandleFunc("/accommodations", controllers.CreateAccommodation).Methods("POST")        // Create a new accommodation
	protected.HandleFunc("/accommodations", controllers.GetAccommodations).Methods("GET")           // Retrieve all accommodations
	protected.HandleFunc("/accommodations/{id}", controllers.GetAccommodation).Methods("GET")       // Retrieve accommodation by ID
	protected.HandleFunc("/accommodations/{id}", controllers.UpdateAccommodation).Methods("PUT")    // Update accommodation
	protected.HandleFunc("/accommodations/{id}", controllers.DeleteAccommodation).Methods("DELETE") // Delete accommodation

	// Booking Management Routes
	protected.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")                   // Create a new booking
	protected.HandleFunc("/bookings/{id}", controllers.GetBooking).Methods("GET")                  // Get a specific booking
	protected.HandleFunc("/bookings/user/{user_id}", controllers.GetBookingsByUser).Methods("GET") // Get all bookings for a user
	protected.HandleFunc("/bookings/{id}/cancel", controllers.CancelBooking).Methods("PUT")        // Cancel a booking

	// Package Booking endpoints
	protected.HandleFunc("/package_bookings", controllers.CreatePackageBooking).Methods("POST")
	protected.HandleFunc("/package_bookings/{booking_id}", controllers.UpdatePackageBooking).Methods("PUT")
	protected.HandleFunc("/package_bookings/{booking_id}", controllers.DeletePackageBooking).Methods("DELETE")

	return router
}
