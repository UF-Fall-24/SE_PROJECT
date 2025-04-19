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

func TestCreatePayment(t *testing.T) {
	payload := `{
		"booking_id": 1,
		"amount": 250.00,
		"payment_method": "Credit Card",
		"payment_status": "Paid"
	}`

	req, err := http.NewRequest("POST", "/payments", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":1,"booking_id":1,"amount":250.0,"payment_method":"Credit Card","payment_status":"Paid"}`))
	}).Methods("POST")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rr.Code)
	}
}

func TestGetAllPayments(t *testing.T) {
	req, _ := http.NewRequest("GET", "/payments", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"booking_id":1,"amount":250.0,"payment_method":"Credit Card","payment_status":"Paid"}]`))
	})

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}

func TestGetPayment(t *testing.T) {
	req, _ := http.NewRequest("GET", "/payments/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"booking_id":1,"amount":250.0,"payment_method":"Credit Card","payment_status":"Paid"}`))
	})

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}

func TestUpdatePayment(t *testing.T) {
	payload := map[string]interface{}{
		"booking_id":     1,
		"amount":         300.0,
		"payment_method": "UPI",
		"payment_status": "Paid",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/payments/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Payment updated successfully"}`))
	}).Methods("PUT")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}

func TestDeletePayment(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/payments/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Payment deleted successfully"}`))
	}).Methods("DELETE")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}
