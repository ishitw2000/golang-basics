package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)

	adminUser := user.NewAdmin(firstName, lastName)

	adminUser.User.OutputUserDetails()

	if err == nil {
		appUser.OutputUserDetails()
		appUser.ClearUsername()
		appUser.OutputUserDetails()
	}

	// ... do something awesome with that gathered data!

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
