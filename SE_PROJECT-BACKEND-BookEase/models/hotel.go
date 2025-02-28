package models

import "time"

// Hotel represents the structure for a hotel record.
type Hotel struct {
	ID          int       `json:"id"`
	HotelName   string    `json:"hotel_name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	RoomType    string    `json:"room_type"`  // e.g., Deluxe, Standard, Suite
	RoomPrice   float64   `json:"room_price"` // Price corresponding to the room type
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
