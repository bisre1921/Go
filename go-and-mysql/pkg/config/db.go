package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// DB is a global variable to store the database connection
var DB *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	var err error

	// MySQL connection string (DSN: Data Source Name)
	dsn := "root:11341149Tedo!@tcp(127.0.0.1:3306)/booking_db"

	// Open a new connection to the database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Check if the connection is actually working
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to MySQL successfully!")
}
