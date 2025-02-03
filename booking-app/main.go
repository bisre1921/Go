package main

import (
	"fmt"
)

const totalTickets int = 50

var conferenceName = "ConferenceName"

var availableTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {

	greetUser()

	for availableTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

		if !isValidName {
			fmt.Println("Please enter a valid name")
			continue
		}

		if !isValidEmail {
			fmt.Println("Please enter a valid email")
			continue
		}

		if !isValidTickets {
			fmt.Printf("Sorry we only have %v tickets available\n", availableTickets)
			continue
		}

		bookTicket(userTickets, firstName, lastName, email)

		printFirstNames()

		var noTicketsRemaining bool = availableTickets == 0
		if noTicketsRemaining {
			fmt.Println("All tickets are sold out")
			break
		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to the %v conference booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available\n", totalTickets, availableTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	fmt.Printf("The first names of bookings are %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	availableTickets = availableTickets - userTickets

	var userData = userData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets for the conference. you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Now we have %v tickets available\n", availableTickets)
}
