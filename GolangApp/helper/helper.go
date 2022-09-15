package helper

import "strings"

func ValidateUserInput(userName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
