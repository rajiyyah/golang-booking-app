package main

import "fmt"

// important notes
// variable names and imported packages must be used to avoid errors
// Paid attention to , when print something!

func main(){
	// fmt.Println("Hello World")

	// list of variables
	var conferenceName = "Go Booking"
	const conferenceTickets int = 50
	var remainingTickets uint= 50
	bookings := []string{}

	
	fmt.Printf("Welcome to %v application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets) // variables names should be in correct order.
	fmt.Println("get your tickets here to attend.")

	for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//User input
		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email : ")
		fmt.Scan(&email)

		fmt.Println("Enter number of ticket that you want : ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		//using slice instead of array
		bookings = append(bookings, firstName + " " + lastName) 

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)

	}

	

}
	