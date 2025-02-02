package middleware

import (
    "book-ease-backend/utils"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v4"
)

func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        // Split the header to get the token
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]

        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return utils.GetJWTKey(), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}
