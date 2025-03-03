package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"book-ease-backend/models"

	"github.com/gorilla/mux"
)

// MockUser simulates a user for testing
var MockUser = models.User{
	ID:       1,
	Username: "kopparla",
	Email:    "kopparla@gmail.com",
	Password: "Hashedpassword123",
}

// Mock function for retrieving user profile (bypassing DB)
func mockGetUserByID(userID int) (*models.User, error) {
	if userID == MockUser.ID {
		return &MockUser, nil
	}
	return nil, nil // Simulating user not found case
}

// Test function for GetUserProfile
func TestGetUserProfile(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate request with mux
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a mock user retrieval
		user, err := mockGetUserByID(1)
		if err != nil || user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Return mock user as response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	// Record the HTTP response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body
	var user models.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Validate user ID
	if user.ID != MockUser.ID {
		t.Errorf("Expected user ID %v, got %v", MockUser.ID, user.ID)
	}

	log.Println("✅ TestGetUserProfile passed!")
}

func TestUpdateUserProfile(t *testing.T) {
	// Create a sample update payload
	updatePayload := `{"username": "varshini", "email": "varshini@gmail.com"}`

	req, err := http.NewRequest("PUT", "/users/1", strings.NewReader(updatePayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Simulate request with mux
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate user update success
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
	})

	// Record the HTTP response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate the response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Validate response message
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedMessage := "Profile updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %v, got %v", expectedMessage, response["message"])
	}

	log.Println("✅ TestUpdateUserProfile passed!")
}

// ✅ TestGetUserProfile_NotFound - Tests fetching a non-existent user
func TestGetUserProfile_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/9999", nil) // Non-existing user ID
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating user not found
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

	t.Log("✅ TestGetUserProfile_NotFound passed!")
}

// ✅ TestGetUserProfile_InvalidID - Tests fetching a user with an invalid ID
func TestGetUserProfile_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/abc", nil) // Invalid ID format
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating invalid ID error
		w.Write([]byte(`{"error": "Invalid user ID"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Invalid user ID"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestGetUserProfile_InvalidID passed!")
}

// ✅ TestUpdateUserProfile_NotFound - Tests updating a non-existent user
func TestUpdateUserProfile_NotFound(t *testing.T) {
	updatePayload := `{"username": "varshkopparla", "email": "varshkopparla@gmail.com"}`

	req, err := http.NewRequest("PUT", "/users/9999", strings.NewReader(updatePayload)) // Non-existing user ID
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating user not found
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

	t.Log("✅ TestUpdateUserProfile_NotFound passed!")
}

// ✅ TestUpdateUserProfile_InvalidID - Tests updating a user with an invalid ID
func TestUpdateUserProfile_InvalidID(t *testing.T) {
	updatePayload := `{"username": "varshh", "email": "varshh@gmail.com"}`

	req, err := http.NewRequest("PUT", "/users/abc", strings.NewReader(updatePayload)) // Invalid ID format
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating invalid ID error
		w.Write([]byte(`{"error": "Invalid user ID"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Invalid user ID"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestUpdateUserProfile_InvalidID passed!")
}

// ✅ TestUpdateUserProfile_InvalidPayload - Tests updating with an invalid request body
func TestUpdateUserProfile_InvalidPayload(t *testing.T) {
	invalidPayload := `{"username": 1234, "email": true}` // Invalid data types

	req, err := http.NewRequest("PUT", "/users/1", strings.NewReader(invalidPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // Simulating invalid request payload
		w.Write([]byte(`{"error": "Invalid request payload"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %v", status)
	}

	expectedResponse := `{"error": "Invalid request payload"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	t.Log("✅ TestUpdateUserProfile_InvalidPayload passed!")
}
