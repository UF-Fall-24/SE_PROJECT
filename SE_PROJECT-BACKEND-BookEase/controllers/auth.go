package controllers

import (
    "book-ease-backend/models"
    "book-ease-backend/utils"
    "encoding/json"
    "log"
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        log.Println("❌ Error decoding request body:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    log.Printf("📝 Registration attempt for username: %s, email: %s", user.Username, user.Email)

    // Validate password
    if err := utils.ValidatePassword(user.Password); err != nil {
        log.Println("❌ Password validation failed:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        log.Println("❌ Error hashing password:", err)
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    // Create user
    err = user.Create()
    if err != nil {
        log.Println("❌ Error creating user:", err)
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    log.Println("✅ User registered successfully:", user.Email)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Decode the JSON request body
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Retrieve user by email
    if err := user.GetByEmail(input.Email); err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    
    // Check if the password is correct
    if err := utils.CheckPassword(input.Password, user.Password); err != nil {
        // In the Login function
log.Printf("🔍 Checking password for user: %s", user.Email)
if err := utils.CheckPassword(input.Password, user.Password); err != nil {
    log.Printf("❌ Password incorrect for user: %s", user.Email)
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
}
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    // Respond with the token
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}


func GetDashboard(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to your dashboard!"})
}
