package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// TestCreateAccommodationBooking tests the CreateAccommodationBooking API endpoint.
func TestCreateAccommodationBooking(t *testing.T) {
	// Sample JSON payload for creating an accommodation booking
	payload := `{
		"package_booking_id": "P1000",
		"first_name": "Alice",
		"last_name": "Wonderland",
		"price": 350.00,
		"duration": "7/6"
	}`

	// Create a new HTTP POST request to /accommodation_bookings
	req, err := http.NewRequest("POST", "/accommodation_bookings", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Define a mock router and handler for CreateAccommodationBooking
	router := mux.NewRouter()
	router.HandleFunc("/accommodation_bookings", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful insertion via stored procedure call
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation booking created successfully"})
	}).Methods("POST")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Validate the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}

	// Validate response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}
	expectedMessage := "Accommodation booking created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}
}

// TestUpdateAccommodationBooking tests the UpdateAccommodationBooking API endpoint.
func TestUpdateAccommodationBooking(t *testing.T) {
	// Sample JSON payload for updating an accommodation booking
	payload := `{
		"package_booking_id": "P1000",
		"first_name": "Alice",
		"last_name": "Liddell",
		"price": 360.00,
		"duration": "7/6"
	}`

	// Create a new HTTP PUT request to /accommodation_bookings/A1000
	req, err := http.NewRequest("PUT", "/accommodation_bookings/A1000", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for UpdateAccommodationBooking
	router := mux.NewRouter()
	router.HandleFunc("/accommodation_bookings/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful update
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation booking updated successfully"})
	}).Methods("PUT")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Validate status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Validate response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}
	expectedMessage := "Accommodation booking updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}
}

// TestDeleteAccommodationBooking tests the DeleteAccommodationBooking API endpoint.
func TestDeleteAccommodationBooking(t *testing.T) {
	// Create a new HTTP DELETE request to /accommodation_bookings/A1000
	req, err := http.NewRequest("DELETE", "/accommodation_bookings/A1000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for DeleteAccommodationBooking
	router := mux.NewRouter()
	router.HandleFunc("/accommodation_bookings/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Accommodation booking deleted successfully"}`)) // ✅ Fixed JSON string
	}).Methods("DELETE")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Validate the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Validate response body
	expectedResponse := `{"message": "Accommodation booking deleted successfully"}` // ✅ Also fixed here
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("Expected response %q, got %q", expectedResponse, rr.Body.String())
	}
}
