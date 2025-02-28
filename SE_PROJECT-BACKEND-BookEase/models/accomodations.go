package models

import (
	"book-ease-backend/config"
	"log"
	"time"
)

// Accommodation struct represents the accommodation entity.
type Accommodation struct {
	ID        int       `json:"id"`
	HotelID   int       `json:"hotel_id"`
	RoomType  string    `json:"room_type"`
	CheckIn   time.Time `json:"check_in"`
	CheckOut  time.Time `json:"check_out"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create inserts a new accommodation record into the database.
func (a *Accommodation) Create() error {
	query := `INSERT INTO accommodations (hotel_id, room_type, check_in, check_out, price)
	          VALUES (?, ?, ?, ?, ?)`
	res, err := config.DB.Exec(query, a.HotelID, a.RoomType, a.CheckIn, a.CheckOut, a.Price)
	if err != nil {
		log.Println("❌ Error inserting accommodation:", err)
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println("❌ Error retrieving last insert ID:", err)
		return err
	}
	a.ID = int(lastID)
	log.Println("✅ Accommodation added successfully with ID:", a.ID)
	return nil
}

// GetByID retrieves an accommodation record by its ID.
func (a *Accommodation) GetByID(id int) error {
	query := "SELECT id, hotel_id, room_type, check_in, check_out, price, created_at, updated_at FROM accommodations WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(
		&a.ID, &a.HotelID, &a.RoomType, &a.CheckIn, &a.CheckOut, &a.Price, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		log.Println("❌ Accommodation not found:", err)
		return err
	}
	log.Println("✅ Accommodation retrieved:", a)
	return nil
}

// GetAll retrieves all accommodation records.
func GetAllAccommodations() ([]Accommodation, error) {
	query := "SELECT id, hotel_id, room_type, check_in, check_out, price, created_at, updated_at FROM accommodations"
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("❌ Error fetching accommodations:", err)
		return nil, err
	}
	defer rows.Close()

	var accommodations []Accommodation
	for rows.Next() {
		var a Accommodation
		err := rows.Scan(&a.ID, &a.HotelID, &a.RoomType, &a.CheckIn, &a.CheckOut, &a.Price, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Println("❌ Error scanning accommodation:", err)
			continue
		}
		accommodations = append(accommodations, a)
	}
	log.Println("✅ Retrieved accommodations:", len(accommodations))
	return accommodations, nil
}

// Update updates an accommodation record by its ID.
func (a *Accommodation) Update() error {
	query := "UPDATE accommodations SET hotel_id=?, room_type=?, check_in=?, check_out=?, price=? WHERE id=?"
	_, err := config.DB.Exec(query, a.HotelID, a.RoomType, a.CheckIn, a.CheckOut, a.Price, a.ID)
	if err != nil {
		log.Println("❌ Error updating accommodation:", err)
		return err
	}
	log.Println("✅ Accommodation updated successfully:", a.ID)
	return nil
}

// Delete removes an accommodation record by its ID.
func (a *Accommodation) Delete() error {
	query := "DELETE FROM accommodations WHERE id=?"
	_, err := config.DB.Exec(query, a.ID)
	if err != nil {
		log.Println("❌ Error deleting accommodation:", err)
		return err
	}
	log.Println("✅ Accommodation deleted successfully:", a.ID)
	return nil
}
