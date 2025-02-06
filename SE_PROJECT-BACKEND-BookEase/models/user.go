package models

import "book-ease-backend/config"
import "log"

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (u *User) Create() error {
    query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
    log.Printf("Executing query: %s", query)
    log.Printf("With values: username=%s, email=%s, password=%s", u.Username, u.Email, u.Password)
    _, err := config.DB.Exec(query, u.Username, u.Email, u.Password)
    return err
}

func (u *User) GetByEmail(email string) error {
    query := "SELECT id, email, password FROM users WHERE email = ?"
    return config.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password)
}

