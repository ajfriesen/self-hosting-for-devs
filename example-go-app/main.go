package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var greeting string
	err := db.QueryRow("SELECT message FROM greetings WHERE id = 1").Scan(&greeting)
	if err != nil {
		log.Printf("Error fetching greeting: %v", err)
		greeting = "I could not fetch a greeting from the database!" // Fallback greeting
	}
	fmt.Fprintf(w, greeting)
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database connection string from the environment
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING not set in .env file")
	}

	// Connect to the database
	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	// Start the server
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}