package models

import "time"

// Package represents the structure of a travel package.
type Package struct {
	ID                 int       `json:"id"`
	PackageName        string    `json:"package_name"`
	PackageDescription string    `json:"package_description"`
	PackagePrice       float64   `json:"package_price"`
	Days               int       `json:"days"`
	Nights             int       `json:"nights"`
	Location           string    `json:"location"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
