package models

import (
	"book-ease-backend/config"
	"log"
	"time"
)

// Booking struct represents a booking record.
type Booking struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	PackageID      int       `json:"package_id"`
	AccommodationID *int     `json:"accommodation_id,omitempty"`
	VehicleID      *int      `json:"vehicle_id,omitempty"`
	BookingDate    time.Time `json:"booking_date"`
	Status         string    `json:"status"`
}

// Create inserts a new booking into the database.
func (b *Booking) Create() error {
	query := `INSERT INTO bookings (user_id, package_id, accommodation_id, vehicle_id, status)
	          VALUES (?, ?, ?, ?, ?)`
	result, err := config.DB.Exec(query, b.UserID, b.PackageID, b.AccommodationID, b.VehicleID, "Confirmed")
	if err != nil {
		log.Println("❌ Error creating booking:", err)
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("❌ Error retrieving last insert ID:", err)
		return err
	}
	b.ID = int(lastID)
	b.Status = "Confirmed"
	b.BookingDate = time.Now()
	log.Println("✅ Booking created successfully:", b.ID)
	return nil
}

// GetByUser retrieves all bookings for a given user.
func GetBookingsByUser(userID int) ([]Booking, error) {
	query := `SELECT id, user_id, package_id, accommodation_id, vehicle_id, booking_date, status FROM bookings WHERE user_id = ?`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("❌ Error retrieving bookings for user:", err)
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var b Booking
		err := rows.Scan(&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID, &b.VehicleID, &b.BookingDate, &b.Status)
		if err != nil {
			log.Println("❌ Error scanning booking row:", err)
			continue
		}
		bookings = append(bookings, b)
	}
	log.Println("✅ Retrieved bookings for user:", userID)
	return bookings, nil
}

// GetByID retrieves a booking by its ID.
func GetBookingByID(bookingID int) (*Booking, error) {
	query := `SELECT id, user_id, package_id, accommodation_id, vehicle_id, booking_date, status FROM bookings WHERE id = ?`
	var b Booking
	err := config.DB.QueryRow(query, bookingID).Scan(&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID, &b.VehicleID, &b.BookingDate, &b.Status)
	if err != nil {
		log.Println("❌ Error retrieving booking:", err)
		return nil, err
	}
	log.Println("✅ Retrieved booking:", bookingID)
	return &b, nil
}

// Cancel updates the booking status to "Canceled".
func (b *Booking) Cancel() error {
	query := `UPDATE bookings SET status = 'Canceled' WHERE id = ?`
	_, err := config.DB.Exec(query, b.ID)
	if err != nil {
		log.Println("❌ Error canceling booking:", err)
		return err
	}
	b.Status = "Canceled"
	log.Println("✅ Booking canceled successfully:", b.ID)
	return nil
}
