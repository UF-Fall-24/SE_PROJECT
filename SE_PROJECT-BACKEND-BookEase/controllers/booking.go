package controllers

import (
	"book-ease-backend/config"
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBooking inserts a new booking (including optional payment_id).
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := booking.Create(); err != nil {
		log.Println("❌ Error creating booking:", err)
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// GetBookingsByUser fetches bookings for a user, including payment_status.
func GetBookingsByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := `
        SELECT 
            b.id, b.user_id, b.package_id, b.accommodation_id, 
            b.payment_id, p.payment_status,
            b.booking_date, b.status
        FROM bookings b
        LEFT JOIN payments p ON b.payment_id = p.payment_id
        WHERE b.user_id = ?`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("❌ Error querying bookings:", err)
		http.Error(w, "Error retrieving bookings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(
			&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID,
			&b.PaymentID, &b.PaymentStatus,
			&b.BookingDate, &b.Status,
		); err != nil {
			log.Println("❌ Error scanning booking row:", err)
			continue
		}
		bookings = append(bookings, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// GetBooking retrieves a booking by its numeric ID, including payment_status.
func GetBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	query := `
        SELECT 
            b.id, b.user_id, b.package_id, b.accommodation_id, 
            b.payment_id, p.payment_status,
            b.booking_date, b.status
        FROM bookings b
        LEFT JOIN payments p ON b.payment_id = p.payment_id
        WHERE b.id = ?`
	var b models.Booking
	if err := config.DB.QueryRow(query, id).Scan(
		&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID,
		&b.PaymentID, &b.PaymentStatus,
		&b.BookingDate, &b.Status,
	); err != nil {
		log.Println("❌ Error retrieving booking:", err)
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

// CancelBooking sets booking status to "Canceled".
func CancelBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	b, err := models.GetBookingByID(id)
	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	if err := b.Cancel(); err != nil {
		log.Println("❌ Error canceling booking:", err)
		http.Error(w, "Error canceling booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled successfully"})
}
