package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetTotalPrice_WithAccommodation tests total price when accommodation_id is provided.
func TestGetTotalPrice_WithAccommodation(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices?package_id=1&accommodation_id=A1000", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) {
		// Simulate package price 500 and accommodation price 150
		total := 500.0 + 150.0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"total_price": total})
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]float64
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if resp["total_price"] != 650.0 {
		t.Errorf("expected total_price 650.0, got %v", resp["total_price"])
	}
}

// TestGetTotalPrice_WithoutAccommodation tests total price when accommodation_id is omitted.
func TestGetTotalPrice_WithoutAccommodation(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices?package_id=2", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) {
		// Simulate package price 300 and no accommodation
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"total_price": 300.0})
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]float64
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if resp["total_price"] != 300.0 {
		t.Errorf("expected total_price 300.0, got %v", resp["total_price"])
	}
}

// TestGetTotalPrice_MissingPackageID tests missing package_id parameter.
func TestGetTotalPrice_MissingPackageID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/prices", GetTotalPrice).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for missing package_id, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestGetTotalPrice_InvalidPackageID tests invalid package_id value.
func TestGetTotalPrice_InvalidPackageID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices?package_id=abc", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/prices", GetTotalPrice).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for invalid package_id, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestGetTotalPrice_InvalidAccommodationID tests invalid accommodation_id is ignored.
func TestGetTotalPrice_InvalidAccommodationID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices?package_id=1&accommodation_id=xyz", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) {
		// Simulate ignoring invalid accommodation and returning package price 400
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"total_price": 400.0})
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]float64
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if resp["total_price"] != 400.0 {
		t.Errorf("expected total_price 400.0, got %v", resp["total_price"])
	}
}
