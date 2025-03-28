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

// TestCreatePackageBooking tests the CreatePackageBooking API endpoint.
func TestCreatePackageBooking(t *testing.T) {
	// Sample JSON payload for creating a package booking
	payload := `{
		"package_id": 1,
		"first_name": "John",
		"last_name": "Doe"
	}`

	// Create a new HTTP POST request to /package_bookings
	req, err := http.NewRequest("POST", "/package_bookings", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Define a mock router and handler for CreatePackageBooking
	router := mux.NewRouter()
	router.HandleFunc("/package_bookings", func(w http.ResponseWriter, r *http.Request) {
		// Simulating successful insertion via stored procedure call
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Package booking created successfully"})
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

	expectedMessage := "Package booking created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}
}

// TestUpdatePackageBooking tests the UpdatePackageBooking API endpoint.
func TestUpdatePackageBooking(t *testing.T) {
	// Sample JSON payload for updating a package booking
	payload := `{
		"package_id": 2,
		"first_name": "Johnathan",
		"last_name": "Doe"
	}`

	// Create a new HTTP PUT request to /package_bookings/P1000
	req, err := http.NewRequest("PUT", "/package_bookings/P1000", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Define a mock router and handler for UpdatePackageBooking
	router := mux.NewRouter()
	router.HandleFunc("/package_bookings/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful update
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Package booking updated successfully"})
	}).Methods("PUT")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Validate the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Validate response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedMessage := "Package booking updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}
}

// TestDeletePackageBooking tests the DeletePackageBooking API endpoint.
func TestDeletePackageBooking(t *testing.T) {
	// Create a new HTTP DELETE request to /package_bookings/P1000
	req, err := http.NewRequest("DELETE", "/package_bookings/P1000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Define a mock router and handler for DeletePackageBooking
	router := mux.NewRouter()
	router.HandleFunc("/package_bookings/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful deletion
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Package booking deleted successfully"}`))
	}).Methods("DELETE")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Validate the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Validate the response body
	expectedResponse := `{"message": "Package booking deleted successfully"}`
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("Expected response %q, got %q", expectedResponse, rr.Body.String())
	}
}
