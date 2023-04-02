package helpers

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

// Dog implements Says and NumberOfLegs so it now satisfies the Animal interface
func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Gorilla implements Says and NumberOfLegs so it now satisfies the Animal interface
func (g Gorilla) Says() string {
	return "Ooo Oooo"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says())
	fmt.Println("This animal has", a.NumberOfLegs(), "legs")
}
