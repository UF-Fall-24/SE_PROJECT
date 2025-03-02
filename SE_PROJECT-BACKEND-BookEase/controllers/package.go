package controllers

import (
	"book-ease-backend/config"
	"book-ease-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePackage handles inserting a new package into the database.
func CreatePackage(w http.ResponseWriter, r *http.Request) {
	var pkg models.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO packages (package_name, package_description, package_price, days, nights, location)
	          VALUES (?, ?, ?, ?, ?, ?)`
	res, err := config.DB.Exec(query, pkg.PackageName, pkg.PackageDescription, pkg.PackagePrice, pkg.Days, pkg.Nights, pkg.Location)
	if err != nil {
		log.Println("Error inserting package:", err)
		http.Error(w, "Error inserting package", http.StatusInternalServerError)
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println("Error retrieving last insert id:", err)
	}
	pkg.ID = int(lastID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pkg)
}

// GetPackages retrieves all packages.
func GetPackages(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, package_name, package_description, package_price, days, nights, location, created_at, updated_at FROM packages")
	if err != nil {
		log.Println("Error retrieving packages:", err)
		http.Error(w, "Error retrieving packages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var packages []models.Package
	for rows.Next() {
		var pkg models.Package
		err := rows.Scan(&pkg.ID, &pkg.PackageName, &pkg.PackageDescription, &pkg.PackagePrice, &pkg.Days, &pkg.Nights, &pkg.Location, &pkg.CreatedAt, &pkg.UpdatedAt)
		if err != nil {
			log.Println("Error scanning package row:", err)
			continue
		}
		packages = append(packages, pkg)
	}
	json.NewEncoder(w).Encode(packages)
}

// GetPackage retrieves a single package by its ID.
func GetPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid package id", http.StatusBadRequest)
		return
	}

	var pkg models.Package
	query := "SELECT id, package_name, package_description, package_price, days, nights, location, created_at, updated_at FROM packages WHERE id = ?"
	err = config.DB.QueryRow(query, id).Scan(&pkg.ID, &pkg.PackageName, &pkg.PackageDescription, &pkg.PackagePrice, &pkg.Days, &pkg.Nights, &pkg.Location, &pkg.CreatedAt, &pkg.UpdatedAt)
	if err != nil {
		log.Println("Error retrieving package:", err)
		http.Error(w, "Package not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pkg)
}

// UpdatePackage updates an existing package by its ID.
func UpdatePackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid package id", http.StatusBadRequest)
		return
	}

	var pkg models.Package
	if err = json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE packages SET package_name=?, package_description=?, package_price=?, days=?, nights=?, location=? WHERE id=?"
	_, err = config.DB.Exec(query, pkg.PackageName, pkg.PackageDescription, pkg.PackagePrice, pkg.Days, pkg.Nights, pkg.Location, id)
	if err != nil {
		log.Println("Error updating package:", err)
		http.Error(w, "Error updating package", http.StatusInternalServerError)
		return
	}
	pkg.ID = id
	json.NewEncoder(w).Encode(pkg)
}

// DeletePackage deletes a package by its ID.
func DeletePackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid package id", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM packages WHERE id=?"
	_, err = config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting package:", err)
		http.Error(w, "Error deleting package", http.StatusInternalServerError)
		return
	}
	// Log and respond that the package was deleted
	log.Println("Package was deleted")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Package was deleted"})

}
