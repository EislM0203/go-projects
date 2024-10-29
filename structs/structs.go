package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {

	appUser, err := user.NewUser(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	)
	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin(
		"test@test.com",
		"123456",
	)

	admin.OutputUserDetails()
	admin.DeleteUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.DeleteUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
