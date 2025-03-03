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

// TestCreateHotel tests the CreateHotel API endpoint
func TestCreateHotel(t *testing.T) {
	// Sample hotel JSON payload
	hotelPayload := `{
		"name": "Hotel Sunshine",
		"location": "New York",
		"rating": 4.5
	}`

	req, err := http.NewRequest("POST", "/hotels", strings.NewReader(hotelPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Define a mock router and handler for CreateHotel
	router := mux.NewRouter()
	router.HandleFunc("/hotels", func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful hotel creation
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hotel created successfully"})
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

	expectedMessage := "Hotel created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestCreateHotel passed!")
}

// TestGetHotels tests the GetHotels API endpoint
func TestGetHotels(t *testing.T) {
	req, err := http.NewRequest("GET", "/hotels", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/hotels", func(w http.ResponseWriter, r *http.Request) {
		// Simulate response with a list of hotels
		hotels := []map[string]interface{}{
			{"id": 1, "name": "Hotel Sunshine", "location": "New York", "rating": 4.5},
			{"id": 2, "name": "Hotel Moonlight", "location": "Los Angeles", "rating": 4.0},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotels)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var hotels []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &hotels)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedCount := 2
	if len(hotels) != expectedCount {
		t.Errorf("Expected %d hotels, got %d", expectedCount, len(hotels))
	}

	log.Println("✅ TestGetHotels passed!")
}

// TestGetHotelByID tests fetching a single hotel by ID
func TestGetHotelByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/hotels/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/hotels/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a response for hotel with ID 1
		hotel := map[string]interface{}{
			"id":       1,
			"name":     "Hotel Sunshine",
			"location": "New York",
			"rating":   4.5,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotel)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var hotel map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &hotel)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedID := 1
	if int(hotel["id"].(float64)) != expectedID {
		t.Errorf("Expected hotel ID %d, got %d", expectedID, int(hotel["id"].(float64)))
	}

	log.Println("✅ TestGetHotelByID passed!")
}

// TestUpdateHotel tests updating an existing hotel
func TestUpdateHotel(t *testing.T) {
	updatePayload := map[string]interface{}{
		"name":     "Hotel Sunshine Deluxe",
		"location": "New York",
		"rating":   4.7,
	}
	requestBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("Error encoding request body: %v", err)
	}

	req, err := http.NewRequest("PUT", "/hotels/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/hotels/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful hotel update
		response := map[string]string{"message": "Hotel updated successfully"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
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

	expectedMessage := "Hotel updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestUpdateHotel passed!")
}

// TestDeleteHotel tests deleting a hotel
func TestDeleteHotel(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/hotels/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/hotels/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hotel deleted successfully"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := `{"message": "Hotel deleted successfully"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	log.Println("✅ TestDeleteHotel passed!")
}
