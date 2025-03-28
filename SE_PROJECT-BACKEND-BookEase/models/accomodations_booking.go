package models

import "time"

// AccommodationBooking represents a record in the accommodation_bookings table.
type AccommodationBooking struct {
	ID                     int        `json:"id"`
	AccommodationBookingID string     `json:"accommodation_booking_id"`
	PackageBookingID       string     `json:"package_booking_id"`
	FirstName              string     `json:"first_name"`
	LastName               string     `json:"last_name"`
	Price                  float64    `json:"price"`
	Duration               string     `json:"duration"` // e.g., "7/6" for days/nights
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}


