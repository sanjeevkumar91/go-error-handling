package main

import (
	"errors"
	"fmt"
	"strings"
)

const Name string = "name"
const Age string = "age"
const Email string = "email"

type User struct {
	Name  string
	Age   int
	Email string
}

func (user *User) IsAdult() bool {
	return user.Age >= 18
}

type ValidationError struct {
	Field string
	Msg   string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", v.Field, v.Msg)
}

func NewUser(name string, age int, email string) (*User, error) {
	if len(strings.TrimSpace(name)) == 0 {
		return nil, &ValidationError{Field: Name, Msg: "name is mandatory"}
	} else if age < 0 || age > 150 {
		return nil, &ValidationError{Field: Age, Msg: "age should be between 0 and 150"}
	} else if !strings.Contains(email, "@") {
		return nil, &ValidationError{Field: Email, Msg: "email should contain @"}
	} else {
		return &User{Name: name, Age: age, Email: email}, nil
	}
}

func main() {
	name := "Ramesh"
	age := 150
	email := "ramesh@gmail.com"

	user, err := NewUser(name, age, email)

	if err != nil {
		// Handle custom error using errors.As
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println(fmt.Errorf("Error in creating user: %w", ve))
		} else {
			// Generic error handler
			fmt.Println("Unhandled error:", err)
		}
	} else {
		fmt.Println("User created:", user.Name)
		fmt.Println("Is Adult:", user.IsAdult())
	}
}
