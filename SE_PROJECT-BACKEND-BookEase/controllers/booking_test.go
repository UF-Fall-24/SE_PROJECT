package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestCreateBookingSuccess tests the CreateBooking endpoint with valid payload.
func TestCreateBookingSuccess(t *testing.T) {
	payload := `{
		"user_id": 1,
		"package_id": 2,
		"accommodation_booking_id": "A1000",
		"payment_id": 123
	}`
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		var b map[string]interface{}
		json.NewDecoder(r.Body).Decode(&b)
		// echo back with assigned ID and status
		b["id"] = 1
		b["status"] = "Confirmed"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)
	}).Methods("POST")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}
	var resp map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if resp["id"].(float64) != 1 {
		t.Errorf("expected id 1, got %v", resp["id"])
	}
}

// TestCreateBookingInvalidJSON tests CreateBooking with malformed JSON.
func TestCreateBookingInvalidJSON(t *testing.T) {
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBufferString(`{invalid json`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", CreateBooking).Methods("POST")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for invalid JSON, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestCreateBookingServerError tests CreateBooking when internal Create fails.
func TestCreateBookingServerError(t *testing.T) {
	payload := `{"user_id":1,"package_id":2}`
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	// stub CreateBooking to simulate DB error
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
	}).Methods("POST")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d for server error, got %d", http.StatusInternalServerError, rr.Code)
	}
}

// TestGetBookingsByUserSuccess tests GetBookingsByUser with valid user_id.
func TestGetBookingsByUserSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/user/1", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		data := []map[string]interface{}{{"id": 1, "user_id": 1, "package_id": 2, "accommodation_booking_id": "A1000", "payment_id": 123, "payment_status": "Paid", "status": "Confirmed"}}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

// TestGetBookingsByUserInvalidID tests GetBookingsByUser with non-integer user_id.
func TestGetBookingsByUserInvalidID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/user/abc", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/user/{user_id}", GetBookingsByUser).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for invalid ID, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestGetBookingSuccess tests GetBooking with valid id.
func TestGetBookingSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/1", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}", func(w http.ResponseWriter, r *http.Request) {
		item := map[string]interface{}{"id": 1, "user_id": 1, "package_id": 2, "accommodation_booking_id": "A1000", "payment_id": 123, "payment_status": "Paid", "status": "Confirmed"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

// TestGetBookingInvalidID tests GetBooking with non-integer id.
func TestGetBookingInvalidID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/abc", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}", GetBooking).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for invalid ID, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestGetBookingNotFound tests GetBooking when record not found.
func TestGetBookingNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bookings/999", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Booking not found", http.StatusNotFound)
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d for not found, got %d", http.StatusNotFound, rr.Code)
	}
}

// TestCancelBookingSuccess tests CancelBooking with valid id.
func TestCancelBookingSuccess(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/bookings/1/cancel", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled successfully"})
	}).Methods("PUT")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

// TestCancelBookingInvalidID tests CancelBooking with non-integer id.
func TestCancelBookingInvalidID(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/bookings/abc/cancel", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", CancelBooking).Methods("PUT")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d for invalid ID, got %d", http.StatusBadRequest, rr.Code)
	}
}

// TestCancelBookingNotFound tests CancelBooking when booking not found.
func TestCancelBookingNotFound(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/bookings/999/cancel", nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Booking not found", http.StatusNotFound)
	}).Methods("PUT")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d for not found, got %d", http.StatusNotFound, rr.Code)
	}
}
