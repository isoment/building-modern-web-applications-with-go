package main

import (
	"fmt"
	"log"

	"github.com/isoment/go-overview/helpers"
)

func main() {
	// Variables and functions
	sum := helpers.Add(2, 2)
	greet, goodbye := helpers.SaySomething("Hi", "See ya")
	log.Println(sum)
	log.Println(greet)
	log.Println(goodbye)

	// Pointers
	myString := "Green"
	helpers.ChangeUsingPointer(&myString)
	fmt.Println(myString)

	// Structs
	user := helpers.User{
		FirstName:   "Bob",
		LastName:    "Smith",
		PhoneNumber: "834-342-2123",
	}
	fmt.Println(user)
	// We can access the PrintFirstName() method on the struct
	var myVar helpers.MyStruct
	myVar.FirstName = "John"
	fmt.Println(myVar.PrintFirstName())
}
