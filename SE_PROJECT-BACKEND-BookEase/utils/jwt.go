package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("se-bookease") // Replace with a secure secret key

// GenerateJWT generates a new JWT token for the user
func GenerateJWT(userID int, email string) (string, error) {
    claims := &jwt.MapClaims{
        "id":    userID,
        "email": email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// GetJWTKey returns the JWT secret key
func GetJWTKey() []byte {
    return jwtKey
}