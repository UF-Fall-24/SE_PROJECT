package controllers

import (
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateAccommodation handles inserting a new accommodation record.
func CreateAccommodation(w http.ResponseWriter, r *http.Request) {
	var accommodation models.Accommodation
	if err := json.NewDecoder(r.Body).Decode(&accommodation); err != nil {
		log.Println("❌ Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := accommodation.Create(); err != nil {
		log.Println("❌ Error creating accommodation:", err)
		http.Error(w, "Error creating accommodation", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Accommodation created successfully:", accommodation.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accommodation)
}

// GetAccommodations retrieves all accommodations.
func GetAccommodations(w http.ResponseWriter, r *http.Request) {
	accommodations, err := models.GetAllAccommodations()
	if err != nil {
		log.Println("❌ Error retrieving accommodations:", err)
		http.Error(w, "Error retrieving accommodations", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Retrieved accommodations count:", len(accommodations))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accommodations)
}

// GetAccommodation retrieves a single accommodation by ID.
func GetAccommodation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("❌ Invalid accommodation ID:", idStr)
		http.Error(w, "Invalid accommodation ID", http.StatusBadRequest)
		return
	}

	var accommodation models.Accommodation
	if err := accommodation.GetByID(id); err != nil {
		log.Println("❌ Accommodation not found:", err)
		http.Error(w, "Accommodation not found", http.StatusNotFound)
		return
	}

	log.Println("✅ Accommodation retrieved:", accommodation.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accommodation)
}

// UpdateAccommodation updates an existing accommodation record.
func UpdateAccommodation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("❌ Invalid accommodation ID:", idStr)
		http.Error(w, "Invalid accommodation ID", http.StatusBadRequest)
		return
	}

	var accommodation models.Accommodation
	if err := json.NewDecoder(r.Body).Decode(&accommodation); err != nil {
		log.Println("❌ Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	accommodation.ID = id

	if err := accommodation.Update(); err != nil {
		log.Println("❌ Error updating accommodation:", err)
		http.Error(w, "Error updating accommodation", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Accommodation updated successfully:", accommodation.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accommodation)
}

// DeleteAccommodation deletes an accommodation record.
func DeleteAccommodation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("❌ Invalid accommodation ID:", idStr)
		http.Error(w, "Invalid accommodation ID", http.StatusBadRequest)
		return
	}

	accommodation := models.Accommodation{ID: id}
	if err := accommodation.Delete(); err != nil {
		log.Println("❌ Error deleting accommodation:", err)
		http.Error(w, "Error deleting accommodation", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Accommodation deleted successfully:", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation deleted successfully"})
}
