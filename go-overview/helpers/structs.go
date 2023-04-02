package helpers

import "time"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type MyStruct struct {
	FirstName string
}

/*
We can add a receiver (m *MyStruct)
This will tie this function to the type of MyStruct and we can access information from MyStruct
*/
func (m *MyStruct) PrintFirstName() string {
	return m.FirstName
}
