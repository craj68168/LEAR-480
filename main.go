package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Bookings System"
	// var name string
	// %v = values
	// %T = types
	// uint = for not negative value -1
	// age:= 26 for another way to declare variable but cannot assign types
	// fmt.Println(&conferenceName) // its the pointer & that print the memory address of variable.
	// fmt.Scan(&name) // it will allow user to input value like alert and get value of that varibale form memory

	// const remainingTickets int = 50
	// var personName = "Raj Chaudhary"
	// fmt.Printf("Welcome to the %v Your available ticket is %v \n", conferenceName, remainingTickets)
	// fmt.Println("Have a good Day.", personName, "Please Visit Again")

	// var bookings = [50]string{"raj","minni"} array having 50 length of data with sobject of string data

	// var bookings []string
	// bookings[0] = firstName + " " + lastName
	// bookings[1] = "nothing"
	// bookings = append(bookings, firstName+" "+lastName) // slice method for array
	// fmt.Printf("Whole Array %v\n", bookings)
	// fmt.Printf("Individual Array %v \n", bookings[0])
	// fmt.Printf("Array Type %T \n", bookings)
	// fmt.Printf("Array length %v \n", len(bookings))

	for {
		var firstName string
		var lastName string
		var email string
		var remainingTickets uint = 50
		var userTicket uint // not allow any negative value
		var bookings []string

		fmt.Println("Enter First Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter Last name")
		fmt.Scan(&lastName)

		fmt.Println(("Enter Email"))
		fmt.Scan(&email)

		fmt.Println("Enter Your Ticket Count")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTicket > 0 && userTicket <= remainingTickets

		if isValidEmail && isValidName && isValidTicket {

			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Your First name is %v and your last name is %v and your ticket count is %v \n", firstName, lastName, userTicket)
			fmt.Printf("Your email is %v \n", email)
			fmt.Printf("Your Remaining ticket is %v \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings { // _ identify unused variables.
				var name = strings.Fields(booking)
				var firstName = name[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("The First name of booking are %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All Ticket are sold out")
				break
			}
		} else {
			fmt.Println("Invalid Input Data")
		}

	}
}
