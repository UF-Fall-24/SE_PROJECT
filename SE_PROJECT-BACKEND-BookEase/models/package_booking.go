package models

import "time"

// PackageBooking represents a record in the package_bookings table.
type PackageBooking struct {
	ID               int       `json:"id"`
	PackageBookingID string    `json:"package_booking_id"`
	PackageID        int       `json:"package_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
