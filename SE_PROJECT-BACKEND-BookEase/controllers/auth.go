package controllers

import (
    "book-ease-backend/models"
	 "book-ease-backend/utils"
    "encoding/json"
    "net/http"
    "log"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Check if username is provided
    if user.Username == "" {
        http.Error(w, "Username is required", http.StatusBadRequest)
        return
    }

    // Validate the password
	if err := utils.ValidatePassword(user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword


    err = user.Create()
    if err != nil {
        log.Printf("Error saving user: %v", err) // Log the error
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials models.User
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Get user by email
    var user models.User
    err = user.GetByEmail(credentials.Email)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Verify password (for simplicity, plaintext password comparison)
    if user.Password != credentials.Password {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Generate JWT
    token, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Respond with the token
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to your dashboard!"})
}
