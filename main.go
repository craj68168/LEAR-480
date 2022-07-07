package main

import (
	"LEAR-480-Go/helper"
	"fmt"
	"strconv"
)

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	for {
		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicket := helper.GetValidateUser(firstName, lastName, email, userTicket, remainingTickets)

		if isValidEmail && isValidName && isValidTicket {

			bookTicket(firstName, lastName, email, userTicket)

			firstNames := getFirstName()
			fmt.Printf("The First name of the booking are %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All Ticket are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is invalid")
			}
			if !isValidEmail {
				fmt.Println("U miss @ in email")
			}
			if !isValidTicket {
				fmt.Println("You input invalid ticket")
			}

		}

	}
}
func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ identify unused variables.
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint // not allow any negative value

	fmt.Println("Enter First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last name")
	fmt.Scan(&lastName)

	fmt.Println(("Enter Email"))
	fmt.Scan(&email)

	fmt.Println("Enter Your Ticket Count")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(firstName string, lastName string, email string, userTicket uint) {
	remainingTickets = remainingTickets - userTicket

	// create a map for user
	var userData = make(map[string]string) // key and value pair with string
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings %v \n", bookings)

	fmt.Printf("Your First name is %v and your last name is %v and your ticket count is %v \n", firstName, lastName, userTicket)
	fmt.Printf("Your email is %v \n", email)
	fmt.Printf("Your Remaining ticket is %v \n", remainingTickets)

}

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

// switch statement
// city := "kathmandu"
// switch city {
// case "chitwan":
// 	// print some
// case "kathmandu":
// // print some
// default:
// 	fmt.Printf("No city found")

// }
