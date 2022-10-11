package main //each package get a folder

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //in go you can return as many values as you want just put them between brackets.
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
/* There are 3 levels of variables scopes
            1- Local : Declaration within function and can be used only in that function , or those declared within a block (e.g for,if-else) and can only be used within it
			2- Package: Declared outside all function, can be used everywhere in the same package, just like bookings variable
			3- Global: Declaration outside all functions and uppercase first letter , Can be used everywhere across all packages */
