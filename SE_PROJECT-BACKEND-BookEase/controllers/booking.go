package controllers

import (
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBooking handles inserting a new booking into the database.
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert booking into database
	err := booking.Create()
	if err != nil {
		log.Println("❌ Error creating booking:", err)
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// GetBookingsByUser retrieves all bookings for a given user.
func GetBookingsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	bookings, err := models.GetBookingsByUser(userID)
	if err != nil {
		log.Println("❌ Error retrieving bookings:", err)
		http.Error(w, "Error retrieving bookings", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookings)
}

// GetBooking retrieves a single booking by its ID.
func GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	booking, err := models.GetBookingByID(bookingID)
	if err != nil {
		log.Println("❌ Error retrieving booking:", err)
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(booking)
}

// CancelBooking cancels a booking by updating its status.
func CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	booking, err := models.GetBookingByID(bookingID)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	err = booking.Cancel()
	if err != nil {
		http.Error(w, "Error canceling booking", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled successfully"})
}
