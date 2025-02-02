package routes

import (
    "book-ease-backend/controllers"
    "github.com/gorilla/mux"
	"book-ease-backend/middleware"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Public routes
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")

    // Protected routes
    protected := router.PathPrefix("/").Subrouter()
    protected.Use(middleware.JWTAuth)

    protected.HandleFunc("/dashboard", controllers.GetDashboard).Methods("GET")

    return router
}
