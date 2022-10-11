package main

import (
	//"booking-app/helper" // use the module which is like a path and add it.
	"fmt"
	"sync"
	"time"
)

// whatever is blaced before the main function is a public variable
const conferenceTickets uint = 50

var conferenceName = "Go Conference" //equivalent of var conferenceName
var remainingTickets uint = 50

// Array definition is var bookings [50]string,this equivalent as well var bookings = [50]string {"Fadia","Souky"} //Go Arrays have a set size it can be declared empty {} or fill it up with few elements {"Fadia","Souky"}
// var bookings = make([]map[string]string, 0) //  add initial size; while arrays are static slices are dynamic
var bookings = make([]UserData, 0)

// bookings []string the [] before string create an empty slice

type UserData struct { //struct is like a lightweight class e.g it doesn't support inhertittance
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//creating a wait group

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	/* for {} is an infinite loop*/
	//for { // we can add conditions to for loops for remainingTickets<50

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) //this is a number of threads that the main thread has to wait for.
		go sendTicket(userTickets, firstName, lastName, email)

		/* here we are creating a seperate thread evrytime we excute this line to send the email instead of blocking the main thread so we can continue a
		accepting new attendance and create a new thread for every newcomer to send the email*/
		/* The problem here is that the main thread can exit without waiting for our sendTivket thread slike if there was no for loop and one execution logic was applied then the app will exit before we even get the chance
		to do our work,a solution is synchronising and telling the main thread to wait here for us*/
		fmt.Printf("These are the first names of all our bookings: %v\n", getFirstNames())

		//var noTicketsRemaining bool= remainingTickets==0// noTicketsRemaining:=remainingTickets==0
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out comeback next year")
			//break
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
	wg.Wait() //this needs to be executed at the end of the main function.
}

//}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string /*after the brackets we specify the return value*/ {
	firstNames := []string{}           //another def for a slice
	for _, booking := range bookings { // for index,booking _ is just a place holder since we won't use index, it is used to indetify unused vars
		// range return the index and value of each elmement (p.s the value can be named anything)
		//firstNames = append(firstNames, booking["firstName"]) //strings.Fields(booking)[0] fields function split a string according to the space and return a slice
		firstNames = append(firstNames, booking.firstName)
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
	//var userData = make(map[string]string) // all keys must have the same type and all values mush be of the same type, we use make to create an empty map.

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	/*userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //formats the uint to a decimal number string.
	*/
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)
	fmt.Printf("Thank you  %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	//sprintf return the string it is formatting
	fmt.Println("####################")
	fmt.Printf("Sending ticket \n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done() //here the thread is saying he is done
}

/*
add in sync sets the number of goroutines to wait for
wait blocks until the WaitGroup counter is 0
Done decrements the waitgroup counter by 1 so this is called by go routines to indicate that it's finished
*/
//a thread in go is called a goroutine , goroutine has the concept of chanels that allows communication between threads
