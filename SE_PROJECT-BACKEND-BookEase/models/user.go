package models

import "book-ease-backend/config"

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (u *User) Create() error {
    query := "INSERT INTO users (email, password) VALUES (?, ?)"
    _, err := config.DB.Exec(query, u.Email, u.Password)
    return err
}

func (u *User) GetByEmail(email string) error {
    query := "SELECT id, email, password FROM users WHERE email = ?"
    return config.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password)
}

