package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// Replace these credentials with your MySQL username, password, and database name
	dsn := "root:Kartik@444@tcp(localhost:3306)/book_ease"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")
	DB = db
}
