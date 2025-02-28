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

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTAuth)

	protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")

	// Package endpoints
	protected.HandleFunc("/packages", controllers.CreatePackage).Methods("POST")
	protected.HandleFunc("/packages/{id}", controllers.UpdatePackage).Methods("PUT")
	protected.HandleFunc("/packages/{id}", controllers.DeletePackage).Methods("DELETE")

	return router
}
