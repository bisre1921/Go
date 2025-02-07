package controllers

import (
	"encoding/json"
	"go-and-mysql/pkg/config"
	"go-and-mysql/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := config.DB.Exec("INSERT INTO bookings (name, email, phone, booking_date) VALUES (?, ?, ?, ?)", book.Name, book.Email, book.Phone, book.BookingDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = int(lastInsertID)
	json.NewEncoder(w).Encode(book)
}

// Get all bookings
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "application/json")
	var books []models.Book
	rows, err := config.DB.Query("SELECT * FROM bookings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Email, &book.Phone, &book.BookingDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

// GetBook gets a single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var book models.Book
	err := config.DB.QueryRow("SELECT * FROM bookings WHERE id = ?", id).Scan(&book.ID, &book.Name, &book.Email, &book.Phone, &book.BookingDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(book)

}
