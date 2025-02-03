package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var isValidTickets bool = userTickets > 0 && userTickets <= availableTickets

	return isValidName, isValidEmail, isValidTickets
}
