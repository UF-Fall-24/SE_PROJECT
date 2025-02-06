package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// Replace these credentials with your MySQL username, password, and database name
	dsn := "root:Vitap@123@tcp(localhost:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// Create database if not exists
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS book_ease")
	if err != nil {
		log.Fatal("Error creating database:", err)
	}

	// Connect to the newly created database
	dsn = "root:Vitap@123@tcp(localhost:3306)/book_ease"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Run migrations
	err = RunMigrations()
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	fmt.Println("Database connected and migrations applied successfully!")

	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Connected to the database!")
	// DB = db
}

func RunMigrations() error {
	script, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(script))
	return err
}
