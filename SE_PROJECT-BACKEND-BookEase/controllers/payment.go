package controllers

import (
	"book-ease-backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var p models.Payment
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := p.Create(); err != nil {
		http.Error(w, "Failed to create payment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p models.Payment
	if err := p.GetByID(id); err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := models.GetAllPayments()
	if err != nil {
		http.Error(w, "Failed to fetch payments", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p models.Payment
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := p.Update(); err != nil {
		http.Error(w, "Failed to update payment", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment updated successfully"})
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	p := models.Payment{ID: id}
	if err := p.Delete(); err != nil {
		http.Error(w, "Failed to delete payment", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment deleted successfully"})
}
