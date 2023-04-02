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
}
