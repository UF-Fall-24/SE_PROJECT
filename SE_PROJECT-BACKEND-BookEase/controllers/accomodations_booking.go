package controllers

import (
	"book-ease-backend/config"
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

// UpdateAccommodationBooking updates an existing accommodation booking using its custom ID.
func UpdateAccommodationBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["booking_id"]

	var ab models.AccommodationBooking
	if err := json.NewDecoder(r.Body).Decode(&ab); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	query := "UPDATE accommodation_bookings SET package_booking_id = ?, first_name = ?, last_name = ?, price = ?, duration = ? WHERE accommodation_booking_id = ?"
	_, err := config.DB.Exec(query, ab.PackageBookingID, ab.FirstName, ab.LastName, ab.Price, ab.Duration, bookingID)
	if err != nil {
		log.Println("Error updating accommodation booking:", err)
		http.Error(w, "Error updating accommodation booking", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation booking updated successfully"})
}

// DeleteAccommodationBooking deletes an accommodation booking using its custom ID.
func DeleteAccommodationBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["booking_id"]

	query := "DELETE FROM accommodation_bookings WHERE accommodation_booking_id = ?"
	_, err := config.DB.Exec(query, bookingID)
	if err != nil {
		log.Println("Error deleting accommodation booking:", err)
		http.Error(w, "Error deleting accommodation booking", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation booking deleted successfully"})
}

