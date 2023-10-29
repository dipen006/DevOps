package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := 50 // same as int32 with this it'll not go to negative
	bookings := []string{} // slices
	// tpyes of var
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferencename is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v Booking App\n", conferenceName)                                                   //formatted output style
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets) // formatted output
	fmt.Println("Get your tickets here to attend")                                                              // normal output

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name and tickets
		// fmt.Scan() // input function
		// userName = "Tom"
		// userTickets = 2
		// fmt.Scan(&userName)

		// fmt.Println(remainingTickets) // returns numbers
		// fmt.Println(&remainingTickets) // returns pointers

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		// isValidCity := city == "Singapore" || "London"
		// isinvalidCity := city != "Singapore" || "London"

		if isValidTicketNumber && isValidName && isValidEmail {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The Whole array is %v\n", bookings)
			fmt.Printf("The first value is %v\n", bookings[0])
			fmt.Printf("Array Type is %T\n", bookings)
			fmt.Printf("Array length is  %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)

			// String = "AB" + "CD"
			// Integers = 5 + 10

			// Array and Slices to store data

			// var bookings = [50]string{"Dipen", "Kunal", "Eddie", "Nana"} // fixed array
			// another inititialization

			// bookings[1] = "Kunal"

			// always remember arrays start from 0th index

			// go detects error at compile time quickly

			// slices
			// mutable arrays means the length of arrays can be changed

			// Dynamic arrrays are callled as slices in Go //

			// slice is abstraction of array
			// slice is more flexible and powerful

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Theses are our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// if remaining ticket is 0 end the prpogram
				fmt.Println("The conference is booked out. Come back next time")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last Name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not correct")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of ticket you entered is invalid")
			}

			fmt.Printf("You input data is invalid !!! Try again.\n")

		}

		// loops
		// we only have for loop in Go

		// infinite loop
		// go to after welcome msg

	} // for loop end

	// for each looop
	// use this to print only first names of the customers

}
