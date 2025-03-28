package controllers

import (
	"book-ease-backend/config"
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
)

// CreateAccommodationBooking calls a stored procedure to insert an accommodation booking.
func CreateAccommodationBooking(w http.ResponseWriter, r *http.Request) {
	var ab models.AccommodationBooking
	if err := json.NewDecoder(r.Body).Decode(&ab); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("CALL InsertAccommodationBooking(?, ?, ?, ?, ?)",
		ab.PackageBookingID, ab.FirstName, ab.LastName, ab.Price, ab.Duration)
	if err != nil {
		log.Println("Error creating accommodation booking:", err)
		http.Error(w, "Error creating accommodation booking", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation booking created successfully"})
}

