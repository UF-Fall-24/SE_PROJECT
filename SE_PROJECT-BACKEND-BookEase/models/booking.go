package models

import (
	"book-ease-backend/config"
	"log"
	"time"
)

// Booking struct represents a booking record, including linked payment status.
type Booking struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	PackageID       int       `json:"package_id"`
	AccommodationID *int      `json:"accommodation_id,omitempty"`
	PaymentID       *int      `json:"payment_id,omitempty"`
	PaymentStatus   *string   `json:"payment_status,omitempty"`
	BookingDate     time.Time `json:"booking_date"`
	Status          string    `json:"status"`
}

// Create inserts a new booking into the database.
func (b *Booking) Create() error {
	query := `INSERT INTO bookings (user_id, package_id, accommodation_id, payment_id, status)
		      VALUES (?, ?, ?, ?, ?)`
	result, err := config.DB.Exec(query, b.UserID, b.PackageID, b.AccommodationID, b.PaymentID, "Confirmed")
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

// GetBookingsByUser retrieves all bookings for a given user, including payment status.
func GetBookingsByUser(userID int) ([]Booking, error) {
	query := `SELECT 
		b.id, b.user_id, b.package_id, b.accommodation_id,
		b.payment_id, p.payment_status,
		b.booking_date, b.status
	  FROM bookings b
	  LEFT JOIN payments p ON b.payment_id = p.payment_id
	  WHERE b.user_id = ?`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("❌ Error retrieving bookings for user:", err)
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var b Booking
		err := rows.Scan(
			&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID,
			&b.PaymentID, &b.PaymentStatus,
			&b.BookingDate, &b.Status,
		)
		if err != nil {
			log.Println("❌ Error scanning booking row:", err)
			continue
		}
		bookings = append(bookings, b)
	}
	log.Println("✅ Retrieved bookings for user:", userID)
	return bookings, nil
}

// GetBookingByID retrieves a booking by its ID, including payment status.
func GetBookingByID(bookingID int) (*Booking, error) {
	query := `SELECT 
		b.id, b.user_id, b.package_id, b.accommodation_id,
		b.payment_id, p.payment_status,
		b.booking_date, b.status
	  FROM bookings b
	  LEFT JOIN payments p ON b.payment_id = p.payment_id
	  WHERE b.id = ?`
	var b Booking
	err := config.DB.QueryRow(query, bookingID).Scan(
		&b.ID, &b.UserID, &b.PackageID, &b.AccommodationID,
		&b.PaymentID, &b.PaymentStatus,
		&b.BookingDate, &b.Status,
	)
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
