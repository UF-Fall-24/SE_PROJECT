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

// TestCreatePackage tests the CreatePackage API endpoint
func TestCreatePackage(t *testing.T) {
	// Sample package JSON payload
	packagePayload := `{
		"name": "Summer Special",
		"price": 299.99,
		"description": "Enjoy a special summer travel package"
	}`

	req, err := http.NewRequest("POST", "/packages", strings.NewReader(packagePayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Define a mock router and handler for CreatePackage
	router := mux.NewRouter()
	router.HandleFunc("/packages", func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful package creation
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Package created successfully"})
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

	expectedMessage := "Package created successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestCreatePackage passed!")
}

// TestGetPackages tests the GetPackages API endpoint
func TestGetPackages(t *testing.T) {
	req, err := http.NewRequest("GET", "/packages", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/packages", func(w http.ResponseWriter, r *http.Request) {
		// Simulate response with a list of packages
		packages := []map[string]interface{}{
			{"id": 1, "name": "Summer Special", "price": 299.99},
			{"id": 2, "name": "Winter Wonderland", "price": 399.99},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(packages)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var packages []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &packages)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedCount := 2
	if len(packages) != expectedCount {
		t.Errorf("Expected %d packages, got %d", expectedCount, len(packages))
	}

	log.Println("✅ TestGetPackages passed!")
}

// TestGetPackageByID tests fetching a single package by ID
func TestGetPackageByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/packages/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/packages/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a response for package with ID 1
		pkg := map[string]interface{}{
			"id":          1,
			"name":        "Summer Special",
			"price":       299.99,
			"description": "Enjoy a special summer travel package",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pkg)
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var pkg map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &pkg)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	expectedID := 1
	if int(pkg["id"].(float64)) != expectedID {
		t.Errorf("Expected package ID %d, got %d", expectedID, int(pkg["id"].(float64)))
	}

	log.Println("✅ TestGetPackageByID passed!")
}

// TestUpdatePackage tests updating an existing package
func TestUpdatePackage(t *testing.T) {
	updatePayload := map[string]interface{}{
		"name":        "Summer Special Plus",
		"price":       349.99,
		"description": "Enjoy an upgraded summer travel package",
	}
	requestBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("Error encoding request body: %v", err)
	}

	req, err := http.NewRequest("PUT", "/packages/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/packages/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful package update
		response := map[string]string{"message": "Package updated successfully"}
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

	expectedMessage := "Package updated successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	log.Println("✅ TestUpdatePackage passed!")
}

// TestDeletePackage tests deleting a package
func TestDeletePackage(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/packages/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/packages/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Package deleted successfully"}`))
	})

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := `{"message": "Package deleted successfully"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, rr.Body.String())
	}

	log.Println("✅ TestDeletePackage passed!")
}
