package main

import "fmt"

func main() {
	// var conferenceName string = "Booking System"
	// var name string
	// %v = values
	// %T = types
	// uint = for not negative value -1
	// age:= 26 for another way to declare variable but cannot assign types
	// fmt.Println(&conferenceName) // its the pointer & that print the memory address of variable.
	// fmt.Scan(&name) // it will allow user to input value like alert and get value of that varibale form memory

	// const conferenceTickets int = 50
	// var personName = "Raj Chaudhary"
	// fmt.Printf("Welcome to the %v Your available ticket is %v \n", conferenceName, conferenceTickets)
	// fmt.Println("Have a good Day.", personName, "Please Visit Again")

	// var booking = [50]string{"raj","minni"} array having 50 length of data with sobject of string data
	var firstName string
	var lastName string
	// var age int

	fmt.Println("Enter First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last name")
	fmt.Scan(&lastName)

	// fmt.Println("Enter Your age")
	// fmt.Scan(&age)
	// fmt.Printf("Your First name is %v and your last name is %v and your age is %v \n", firstName, lastName, age)

	var booking []string
	// booking[0] = firstName + " " + lastName
	// booking[1] = "nothing"
	booking = append(booking, firstName+" "+lastName) // slice method for array
	fmt.Printf("Whole Array %v\n", booking)
	fmt.Printf("Individual Array %v \n", booking[0])
	fmt.Printf("Array Type %T \n", booking)
	fmt.Printf("Array length %v \n", len(booking))
}
