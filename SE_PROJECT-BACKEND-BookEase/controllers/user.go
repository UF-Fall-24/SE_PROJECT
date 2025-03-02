package controllers

import (
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUserProfile retrieves user details by ID
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("🔍 GetUserProfile API hit") // Debug log

	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		log.Println("❌ No 'id' parameter found in request")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("❌ Invalid user ID:", idStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	log.Println("🆔 Fetching user ID:", id)

	var user models.User
	if err := user.GetUserByID(id); err != nil {
		log.Println("❌ Error retrieving user profile:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Println("✅ User profile retrieved successfully:", user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserProfile updates the username or email of a user
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("🔄 UpdateUserProfile API hit")

	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		log.Println("❌ No 'id' parameter found in request")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("❌ Invalid user ID:", idStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.Println("❌ Error decoding request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updateData.ID = id // Assign ID to struct

	if err := updateData.UpdateUser(); err != nil {
		log.Println("❌ Error updating user profile:", err)
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Profile updated successfully for user ID:", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
