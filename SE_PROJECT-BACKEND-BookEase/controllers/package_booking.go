package controllers

import (
	"book-ease-backend/config"
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreatePackageBooking calls a stored procedure to insert a package booking.
func CreatePackageBooking(w http.ResponseWriter, r *http.Request) {
	var pb models.PackageBooking
	if err := json.NewDecoder(r.Body).Decode(&pb); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("CALL InsertPackageBooking(?, ?, ?)", pb.PackageID, pb.FirstName, pb.LastName)
	if err != nil {
		log.Println("Error creating package booking:", err)
		http.Error(w, "Error creating package booking", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Package booking created successfully"})
}

// UpdatePackageBooking updates an existing package booking using its custom ID.
func UpdatePackageBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["booking_id"]

	var pb models.PackageBooking
	if err := json.NewDecoder(r.Body).Decode(&pb); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	query := "UPDATE package_bookings SET package_id = ?, first_name = ?, last_name = ? WHERE package_booking_id = ?"
	_, err := config.DB.Exec(query, pb.PackageID, pb.FirstName, pb.LastName, bookingID)
	if err != nil {
		log.Println("Error updating package booking:", err)
		http.Error(w, "Error updating package booking", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Package booking updated successfully"})
}

// DeletePackageBooking deletes a package booking using its custom ID.
func DeletePackageBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["booking_id"]

	query := "DELETE FROM package_bookings WHERE package_booking_id = ?"
	_, err := config.DB.Exec(query, bookingID)
	if err != nil {
		log.Println("Error deleting package booking:", err)
		http.Error(w, "Error deleting package booking", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Package booking deleted successfully"})
}
