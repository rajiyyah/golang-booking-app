package main

import (
	"fmt"
	"sync"
	"time"
)

// important notes
// variable names and imported packages must be used to avoid errors
// Paid attention to , when print something!

// This is a package level variable where we list all of variables that are accessible to all functions
const conferenceTickets int = 50

var conferenceName = "Go Booking"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var wg = sync.WaitGroup{} //function: to wait for the launched application to finish

func main() {
	// fmt.Println("Hello World")

	// calling function
	greetUsers()

	for {

		// calling function
		firstName, lastName, email, userTickets := getUserInput()

		// calling function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//calling function
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)                                              // increase the number of thread that should be wait
			go sendTicket(userTickets, firstName, lastName, email) // go keyword : starts a new thread/ goroutine

			// calling function print first name
			firstNames := getFirstNames()
			fmt.Printf(" The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our ticket is all booked out. Come back next year.")
				break // to stop the application & skip the remaining program
			}

		} else {
			if !isValidName {
				fmt.Println("Your first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address that you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}

		}
		wg.Wait()

	}

}

func greetUsers() {

	fmt.Printf("Welcome to %v application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets) // variables names should be in correct order.
	fmt.Println("get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ is a blank identifier, to ignore a variable you dont want to use
		firstNames = append(firstNames, booking.firstName) //append(firstNames, booking["firstName"]) --> use [] if use map
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//User input
	fmt.Println("\nEnter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket that you want : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// // Example : create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//create struct
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTickets,
	}

	//using slice instead of array
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() //
}
