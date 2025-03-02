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

// CreateHotel handles inserting a new hotel into the database.
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO hotels (hotel_name, address, city, description, rating, room_type, room_price)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := config.DB.Exec(query, hotel.HotelName, hotel.Address, hotel.City, hotel.Description, hotel.Rating, hotel.RoomType, hotel.RoomPrice)
	if err != nil {
		log.Println("Error inserting hotel:", err)
		http.Error(w, "Error inserting hotel", http.StatusInternalServerError)
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println("Error retrieving last insert id:", err)
	}
	hotel.ID = int(lastID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hotel)
}

// GetHotels retrieves all hotels.
func GetHotels(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, hotel_name, address, city, description, rating, room_type, room_price, created_at, updated_at FROM hotels")
	if err != nil {
		log.Println("Error retrieving hotels:", err)
		http.Error(w, "Error retrieving hotels", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.Address, &hotel.City, &hotel.Description, &hotel.Rating, &hotel.RoomType, &hotel.RoomPrice, &hotel.CreatedAt, &hotel.UpdatedAt)
		if err != nil {
			log.Println("Error scanning hotel row:", err)
			continue
		}
		hotels = append(hotels, hotel)
	}
	json.NewEncoder(w).Encode(hotels)
}

// GetHotel retrieves a single hotel by its ID.
func GetHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid hotel id", http.StatusBadRequest)
		return
	}

	var hotel models.Hotel
	query := "SELECT id, hotel_name, address, city, description, rating, room_type, room_price, created_at, updated_at FROM hotels WHERE id = ?"
	err = config.DB.QueryRow(query, id).Scan(&hotel.ID, &hotel.HotelName, &hotel.Address, &hotel.City, &hotel.Description, &hotel.Rating, &hotel.RoomType, &hotel.RoomPrice, &hotel.CreatedAt, &hotel.UpdatedAt)
	if err != nil {
		log.Println("Error retrieving hotel:", err)
		http.Error(w, "Hotel not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(hotel)
}

// UpdateHotel updates an existing hotel by its ID.
func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid hotel id", http.StatusBadRequest)
		return
	}

	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE hotels SET hotel_name=?, address=?, city=?, description=?, rating=?, room_type=?, room_price=? WHERE id=?"
	_, err = config.DB.Exec(query, hotel.HotelName, hotel.Address, hotel.City, hotel.Description, hotel.Rating, hotel.RoomType, hotel.RoomPrice, id)
	if err != nil {
		log.Println("Error updating hotel:", err)
		http.Error(w, "Error updating hotel", http.StatusInternalServerError)
		return
	}
	hotel.ID = id
	json.NewEncoder(w).Encode(hotel)
}

// DeleteHotel deletes a hotel by its ID.
func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid hotel id", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM hotels WHERE id=?"
	_, err = config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting hotel:", err)
		http.Error(w, "Error deleting hotel", http.StatusInternalServerError)
		return
	}

	log.Println("Hotel was deleted")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hotel was deleted"})
}

// GetHotelsByLocation retrieves hotels filtered by the provided location (city).
func GetHotelsByLocation(w http.ResponseWriter, r *http.Request) {
	// Read the location query parameter
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, "Location parameter is required", http.StatusBadRequest)
		return
	}

	// Query the hotels table for records matching the specified location (city)
	query := "SELECT id, hotel_name, address, city, description, rating, room_type, room_price, created_at, updated_at FROM hotels WHERE city = ?"
	rows, err := config.DB.Query(query, location)
	if err != nil {
		log.Println("Error retrieving hotels by location:", err)
		http.Error(w, "Error retrieving hotels", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.Address, &hotel.City, &hotel.Description, &hotel.Rating, &hotel.RoomType, &hotel.RoomPrice, &hotel.CreatedAt, &hotel.UpdatedAt)
		if err != nil {
			log.Println("Error scanning hotel row:", err)
			continue
		}
		hotels = append(hotels, hotel)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}
