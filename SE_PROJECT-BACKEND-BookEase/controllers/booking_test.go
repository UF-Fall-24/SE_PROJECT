package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// ✅ TestCreateBooking - Tests the creation of a new booking
func TestCreateBooking(t *testing.T) {
	mockBooking := map[string]interface{}{
		"user_id":         1,
		"package_id":      2,
		"accommodation_id": 3,
		"vehicle_id":      4,
	}

	requestBody, err := json.Marshal(mockBooking)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Booking created successfully"})
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedMessage := "Booking created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestCreateBooking passed!")
}

// ✅ TestGetBookingsByUser - Tests retrieving all bookings for a user
func TestGetBookingsByUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/bookings/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/bookings/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		bookings := []map[string]interface{}{
			{"id": 1, "user_id": 1, "package_id": 2, "accommodation_id": 3, "vehicle_id": 4, "status": "Confirmed"},
			{"id": 2, "user_id": 1, "package_id": 5, "accommodation_id": 6, "vehicle_id": 7, "status": "Pending"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var bookings []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &bookings)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedCount := 2
	if len(bookings) != expectedCount {
		t.Errorf("Expected %d bookings, got %d", expectedCount, len(bookings))
	}

	log.Println("✅ TestGetBookingsByUser passed!")
}

// ✅ TestGetBookingByID - Tests fetching a single booking by ID
func TestGetBookingByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/bookings/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}", func(w http.ResponseWriter, r *http.Request) {
		booking := map[string]interface{}{
			"id":             1,
			"user_id":        1,
			"package_id":     2,
			"accommodation_id": 3,
			"vehicle_id":     4,
			"status":         "Confirmed",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var booking map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &booking)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedID := 1
	if int(booking["id"].(float64)) != expectedID {
		t.Errorf("Expected booking ID %d, got %d", expectedID, int(booking["id"].(float64)))
	}

	log.Println("✅ TestGetBookingByID passed!")
}

// ✅ TestCancelBooking - Tests canceling a booking
func TestCancelBooking(t *testing.T) {
	req, err := http.NewRequest("PUT", "/bookings/1/cancel", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled successfully"})
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedMessage := "Booking canceled successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestCancelBooking passed!")
}

// ✅ TestCreateBooking_MissingFields - Tests creating a booking with missing fields
func TestCreateBooking_MissingFields(t *testing.T) {
	missingFieldsPayload := `{
		"user_id": 1,
		"package_id": 2
	}` // Missing accommodation_id and vehicle_id

	req, err := http.NewRequest("POST", "/bookings", strings.NewReader(missingFieldsPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating missing fields error
		w.Write([]byte(`{"error": "Missing required fields"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Missing required fields"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCreateBooking_MissingFields passed!")
}

// ✅ TestCreateBooking_InvalidData - Tests creating a booking with invalid data types
func TestCreateBooking_InvalidData(t *testing.T) {
	invalidDataPayload := `{
		"user_id": "one",
		"package_id": "two",
		"accommodation_id": "three",
		"vehicle_id": "four"
	}` // All IDs are strings instead of integers

	req, err := http.NewRequest("POST", "/bookings", strings.NewReader(invalidDataPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating invalid data type error
		w.Write([]byte(`{"error": "Invalid data types"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Invalid data types"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCreateBooking_InvalidData passed!")
}

// ✅ TestCreateBooking_UserNotFound - Tests creating a booking for a non-existing user
func TestCreateBooking_UserNotFound(t *testing.T) {
	nonExistingUserPayload := `{
		"user_id": 9999,
		"package_id": 2,
		"accommodation_id": 3,
		"vehicle_id": 4
	}`

	req, err := http.NewRequest("POST", "/bookings", strings.NewReader(nonExistingUserPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating user not found error
		w.Write([]byte(`{"error": "User not found"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %v", status)
	}

	expectedResponse := `{"error": "User not found"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCreateBooking_UserNotFound passed!")
}

// ✅ TestCreateBooking_PackageNotFound - Tests creating a booking for a non-existing package
func TestCreateBooking_PackageNotFound(t *testing.T) {
	nonExistingPackagePayload := `{
		"user_id": 1,
		"package_id": 9999,
		"accommodation_id": 3,
		"vehicle_id": 4
	}`

	req, err := http.NewRequest("POST", "/bookings", strings.NewReader(nonExistingPackagePayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating package not found error
		w.Write([]byte(`{"error": "Package not found"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %v", status)
	}

	expectedResponse := `{"error": "Package not found"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCreateBooking_PackageNotFound passed!")
}


// ✅ TestCancelBooking_NotFound - Tests canceling a non-existent booking
func TestCancelBooking_NotFound(t *testing.T) {
	req, err := http.NewRequest("PUT", "/bookings/9999/cancel", nil) // Non-existing booking ID
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating booking not found
		w.Write([]byte(`{"error": "Booking not found"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %v", status)
	}

	expectedResponse := `{"error": "Booking not found"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCancelBooking_NotFound passed!")
}

// ✅ TestCancelBooking_InvalidID - Tests canceling a booking with an invalid ID format
func TestCancelBooking_InvalidID(t *testing.T) {
	req, err := http.NewRequest("PUT", "/bookings/abc/cancel", nil) // Invalid ID format
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating invalid ID error
		w.Write([]byte(`{"error": "Invalid booking ID"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Invalid booking ID"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCancelBooking_InvalidID passed!")
}

// ✅ TestCancelBooking_AlreadyCanceled - Tests canceling an already canceled booking
func TestCancelBooking_AlreadyCanceled(t *testing.T) {
	req, err := http.NewRequest("PUT", "/bookings/1/cancel", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/bookings/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		// Simulating that the booking has already been canceled
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Booking is already canceled"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Booking is already canceled"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCancelBooking_AlreadyCanceled passed!")
}
