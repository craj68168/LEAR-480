package helper

import "strings"

func GetValidateUser(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}

// to use helper.go and its function to main.go run go run main.go helper.go
// to run multiple files go run .
