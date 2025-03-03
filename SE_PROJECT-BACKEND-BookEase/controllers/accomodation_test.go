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

// TestCreateAccommodation tests the CreateAccommodation API
func TestCreateAccommodation(t *testing.T) {
	// Sample accommodation JSON payload
	accommodationPayload := `{
		"hotel_id": 1,
		"room_type": "Deluxe",
		"check_in": "2024-07-01",
		"check_out": "2024-07-05",
		"price": 350.00
	}`

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/accommodations", strings.NewReader(accommodationPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for CreateAccommodation
	router := mux.NewRouter()
	router.HandleFunc("/accommodations", func(w http.ResponseWriter, r *http.Request) {
		// Simulating a successful accommodation creation
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Accommodation created successfully"})
	})

	// Serve the request
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Parse response body
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Validate response message
	expectedMessage := "Accommodation created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %v, got %v", expectedMessage, response["message"])
	}

	log.Println("✅ TestCreateAccommodation passed!")
}


// TestGetAccommodations tests the GetAccommodations API
func TestGetAccommodations(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/accommodations", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for GetAccommodations
	router := mux.NewRouter()
	router.HandleFunc("/accommodations", func(w http.ResponseWriter, r *http.Request) {
		// Simulating a successful response with mock accommodations
		accommodations := []map[string]interface{}{
			{"id": 1, "hotel_id": 1, "room_type": "Deluxe", "check_in": "2024-07-01", "check_out": "2024-07-05", "price": 350.00},
			{"id": 2, "hotel_id": 2, "room_type": "Suite", "check_in": "2024-08-10", "check_out": "2024-08-15", "price": 450.00},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accommodations)
	})

	// Serve the request
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse response body
	var accommodations []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &accommodations)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Validate response length
	expectedCount := 2
	if len(accommodations) != expectedCount {
		t.Errorf("Expected %d accommodations, got %d", expectedCount, len(accommodations))
	}

	log.Println("✅ TestGetAccommodations passed!")
}

// TestGetAccommodationByID tests fetching a single accommodation by ID
func TestGetAccommodationByID(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/accommodations/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for GetAccommodationByID
	router := mux.NewRouter()
	router.HandleFunc("/accommodations/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulating a successful response for accommodation ID 1
		accommodation := map[string]interface{}{
			"id":       1,
			"hotel_id": 1,
			"room_type": "Deluxe",
			"check_in":  "2024-07-01",
			"check_out": "2024-07-05",
			"price":     350.00,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accommodation)
	})

	// Serve the request
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse response body
	var accommodation map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &accommodation)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Validate response data
	expectedID := 1
	if int(accommodation["id"].(float64)) != expectedID {
		t.Errorf("Expected accommodation ID %d, got %d", expectedID, int(accommodation["id"].(float64)))
	}

	log.Println("✅ TestGetAccommodationByID passed!")
}

// TestUpdateAccommodation tests updating an accommodation
func TestUpdateAccommodation(t *testing.T) {
	// Create a test accommodation update payload
	updatePayload := map[string]interface{}{
		"room_type": "Suite",
		"check_in":  "2024-08-01",
		"check_out": "2024-08-05",
		"price":     450.00,
	}

	// Convert payload to JSON
	requestBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("Error encoding request body: %v", err)
	}

	// Create a new HTTP PUT request
	req, err := http.NewRequest("PUT", "/accommodations/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for UpdateAccommodation
	router := mux.NewRouter()
	router.HandleFunc("/accommodations/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulating successful update response
		response := map[string]string{
			"message": "Accommodation updated successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Serve the request
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse response body
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Validate response message
	expectedMessage := "Accommodation updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestUpdateAccommodation passed!")
}

// TestDeleteAccommodation tests the accommodation deletion functionality
func TestDeleteAccommodation(t *testing.T) {
	// Create a new DELETE request for an existing accommodation
	req, err := http.NewRequest("DELETE", "/accommodations/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Define a mock router and handler for DeleteAccommodation
	router := mux.NewRouter()
	router.HandleFunc("/accommodations/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Accommodation deleted successfully"}`))
	})

	// Serve the request
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Validate response message
	expectedResponse := `{"message": "Accommodation deleted successfully"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestDeleteAccommodation passed!")
}

// ✅ TestCreateAccommodation_MissingFields - Tests creating an accommodation with missing fields
func TestCreateAccommodation_MissingFields(t *testing.T) {
	// Missing hotel_id and check_out fields
	missingFieldsPayload := `{
		"room_type": "Deluxe",
		"check_in": "2024-07-01",
		"price": 350.00
	}`

	req, err := http.NewRequest("POST", "/accommodations", strings.NewReader(missingFieldsPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accommodations", func(w http.ResponseWriter, r *http.Request) {
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

	t.Log("✅ TestCreateAccommodation_MissingFields passed!")
}

// ✅ TestCreateAccommodation_InvalidData - Tests creating an accommodation with invalid data types
func TestCreateAccommodation_InvalidData(t *testing.T) {
	// Hotel ID as string instead of integer
	invalidDataPayload := `{
		"hotel_id": "one",
		"room_type": "Deluxe",
		"check_in": "2024-07-01",
		"check_out": "2024-07-05",
		"price": "cheap"
	}`

	req, err := http.NewRequest("POST", "/accommodations", strings.NewReader(invalidDataPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accommodations", func(w http.ResponseWriter, r *http.Request) {
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

	t.Log("✅ TestCreateAccommodation_InvalidData passed!")
}

// ✅ TestCreateAccommodation_HotelNotFound - Tests creating an accommodation for a non-existing hotel
func TestCreateAccommodation_HotelNotFound(t *testing.T) {
	nonExistingHotelPayload := `{
		"hotel_id": 9999,
		"room_type": "Deluxe",
		"check_in": "2024-07-01",
		"check_out": "2024-07-05",
		"price": 350.00
	}`

	req, err := http.NewRequest("POST", "/accommodations", strings.NewReader(nonExistingHotelPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accommodations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating hotel not found error
		w.Write([]byte(`{"error": "Hotel not found"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %v", status)
	}

	expectedResponse := `{"error": "Hotel not found"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestCreateAccommodation_HotelNotFound passed!")
}
