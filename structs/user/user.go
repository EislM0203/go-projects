package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Printf("First name: %s, Last name: %s, Birthdate: %s, Created at: %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) DeleteUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser (firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("no empty strings allowed")
	}


	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin (email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "admin", 
			lastName: "admin", 
			birthdate: "---", 
			createdAt: time.Now(),
		},
	}
}