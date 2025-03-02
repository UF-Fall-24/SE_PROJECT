package models

import (
	"book-ease-backend/config"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (u *User) Create() error {
    query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
    log.Printf("Executing query: %s", query)
    log.Printf("With values: username=%s, email=%s", u.Username, u.Email)
    _, err := config.DB.Exec(query, u.Username, strings.ToLower(u.Email), u.Password)
    return err
}

func (u *User) GetByEmail(email string) error {
    query := "SELECT id, email, password FROM users WHERE email = ?"
    
    log.Println("🔍 Running SQL Query:", query)
    log.Println("🔍 Searching for email:", email)

    err := config.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password)
    if err != nil {
        log.Println("❌ User not found in database:", err)
        return err
    }

    log.Println("✅ User found:", u.Email)
    return nil
}

// GetUserByID retrieves user details by ID.
func (u *User) GetUserByID(userID int) error {
	query := "SELECT id, username, email FROM users WHERE id = ?"
	log.Println("🔍 Fetching user by ID:", userID)

	err := config.DB.QueryRow(query, userID).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("❌ User not found:", userID)
			return fmt.Errorf("user not found")
		}
		log.Println("❌ Database error:", err)
		return err
	}

	log.Println("✅ User found:", u.Username)
	return nil
}

// UpdateUser updates user details (excluding password) in the database.
func (u *User) UpdateUser() error {
	query := "UPDATE users SET username = ?, email = ? WHERE id = ?"

	log.Println("🔄 Updating user:", u.Username)

	_, err := config.DB.Exec(query, u.Username, strings.ToLower(u.Email), u.ID)
	if err != nil {
		log.Println("❌ Error updating user:", err)
		return err
	}

	log.Println("✅ User updated successfully")
	return nil
}