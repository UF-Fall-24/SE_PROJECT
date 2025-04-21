package controllers

import (
	"book-ease-backend/config"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetTotalPrice returns the sum of package price and (optional) accommodation price.
func GetTotalPrice(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// 1) package_id is required
	pkgIDStr := q.Get("package_id")
	if pkgIDStr == "" {
		http.Error(w, "package_id query parameter is required", http.StatusBadRequest)
		return
	}
	pkgID, err := strconv.Atoi(pkgIDStr)
	if err != nil {
		http.Error(w, "invalid package_id", http.StatusBadRequest)
		return
	}

	// 2) fetch package_price
	var pkgPrice float64
	err = config.DB.QueryRow(
		"SELECT package_price FROM packages WHERE id = ?", pkgID,
	).Scan(&pkgPrice)
	if err != nil {
		log.Println("❌ Error fetching package price:", err)
		http.Error(w, "package not found", http.StatusNotFound)
		return
	}

	total := pkgPrice

	// 3) if accommodation_id provided, fetch and add
	if accIDStr := q.Get("accommodation_id"); accIDStr != "" {

		var accPrice float64
		if err := config.DB.QueryRow(
			"SELECT price FROM accommodation_bookings WHERE accommodation_booking_id = ?", accIDStr,
		).Scan(&accPrice); err == nil {
			total += accPrice
		}
		// if error, we simply ignore accommodation price

	}

	// 4) return only the total_price
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_price": total})
}
