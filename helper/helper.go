package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, emailAdress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAdress, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}
