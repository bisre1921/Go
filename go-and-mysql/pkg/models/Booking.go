package models

// Booking struct represents a booking
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	BookingDate string `json:"booking_date"`
}
