package models

import (
	"book-ease-backend/config"
	"log"
	"time"
)

type Payment struct {
	ID            int       `json:"id"`
	BookingID     int       `json:"booking_id"`
	Amount        float64   `json:"amount"`
	Method        string    `json:"payment_method"`
	Status        string    `json:"payment_status"`
	PaymentDate   time.Time `json:"payment_date"`
}

// Create inserts a new payment record.
func (p *Payment) Create() error {
	query := `INSERT INTO payments (booking_id, amount, payment_method, payment_status)
	          VALUES (?, ?, ?, ?)`
	res, err := config.DB.Exec(query, p.BookingID, p.Amount, p.Method, p.Status)
	if err != nil {
		log.Println("❌ Error inserting payment:", err)
		return err
	}
	id, _ := res.LastInsertId()
	p.ID = int(id)
	return nil
}

// GetByID fetches a payment by ID.
func (p *Payment) GetByID(id int) error {
	query := `SELECT id, booking_id, amount, payment_method, payment_status, payment_date
	          FROM payments WHERE id = ?`
	return config.DB.QueryRow(query).Scan(&p.ID, &p.BookingID, &p.Amount, &p.Method, &p.Status, &p.PaymentDate)
}

// GetAll fetches all payments.
func GetAllPayments() ([]Payment, error) {
	query := `SELECT id, booking_id, amount, payment_method, payment_status, payment_date FROM payments`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.ID, &p.BookingID, &p.Amount, &p.Method, &p.Status, &p.PaymentDate); err != nil {
			continue
		}
		payments = append(payments, p)
	}
	return payments, nil
}

// Update modifies an existing payment.
func (p *Payment) Update() error {
	query := `UPDATE payments SET amount=?, payment_method=?, payment_status=? WHERE id=?`
	_, err := config.DB.Exec(query, p.Amount, p.Method, p.Status, p.ID)
	return err
}

// Delete removes a payment by ID.
func (p *Payment) Delete() error {
	query := `DELETE FROM payments WHERE id = ?`
	_, err := config.DB.Exec(query, p.ID)
	return err
}
