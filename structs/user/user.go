package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastname  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastname:  "ADMIN",
			birthdate: "-----",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastname, u.birthdate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastname = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastname:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
