package main

import (
	"booking-app/helper" // use the module which is like a path and add it.
	"fmt"
	"strings"
)

// whatever is blaced before the main function is a public variable
const conferenceTickets uint = 50

var conferenceName = "Go Conference" //equivalent of var conferenceName
var remainingTickets uint = 50

// Array definition is var bookings [50]string,this equivalent as well var bookings = [50]string {"Fadia","Souky"} //Go Arrays have a set size it can be declared empty {} or fill it up with few elements {"Fadia","Souky"}
var bookings []string // while arrays are static slices are dynamic
func main() {

	greetUser()

	/* for {} is an infinite loop*/
	for { // we can add conditions to for loops for remainingTickets<50

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			fmt.Printf("These are the first names of all our bookings: %v\n", getFirstNames())

			//var noTicketsRemaining bool= remainingTickets==0// noTicketsRemaining:=remainingTickets==0
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out comeback next year")
				break
			}
		} else if userTickets == remainingTickets { //here you can have as much else if as you want

		} else {

			if !isValidName {
				fmt.Println("The first or last name you entered is too short ")

			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalide ")

			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets  you entered is invalide ")

			}

		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string /*after the brackets we specify the return value*/ {
	firstNames := []string{}           //another def for a slice
	for _, booking := range bookings { // for index,booking _ is just a place holder since we won't use index, it is used to indetify unused vars
		// range return the index and value of each elmement (p.s the value can be named anything)
		firstNames = append(firstNames, strings.Fields(booking)[0]) //fields function split a string according to the space and return a slice

	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) // fmt.println(userName) will print the value of the variable while fmt.Println(&userName) will print the pointer: where the value is stored.

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
