package models

import (
    "book-ease-backend/config"
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
