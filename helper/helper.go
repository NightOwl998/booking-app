package helper //each package get a folder

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //in go you can return as many values as you want just put them between brackets.
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

/* Now that we have added the helper.go file to run we need to specify both
     the main file and helper bu "go run main.go helper.go" or run simply "go run ."
	 here we are specifying a folder to run from. */
/* code in the same package has acess to anything global varaible and functions, anything outside a function is acessible in the same package and all*/
/* to make a function avalaible in other package we have to export it by capatilising the first letter of the name*/
/* You can also export variables,function,constants and types from other packages*/
