package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

// We can define a struct for the JSON
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

/*
Converting JSON into a struct
*/
func UnmarshallJSON() {
	myJson := `
	[
		{
			"first_name": "Bill",
			"last_name": "Smith",
			"has_dog": false
		},
		{
			"first_name": "Sally",
			"last_name": "Doe",
			"has_dog": true
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
}

/*
Converting a slice of structs into JSON
*/
func MarshallJSON() {
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Stan"
	m1.LastName = "Jones"
	m1.HasDog = true

	var m2 Person
	m2.FirstName = "Adam"
	m2.LastName = "West"
	m2.HasDog = false

	mySlice = append(mySlice, m1, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshalling json", err)
	}

	fmt.Println(string(newJson))
}
