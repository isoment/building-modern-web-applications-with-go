package helpers

import "fmt"

type UserAccount struct {
	FirstName string
	LastName  string
}

/*
Create a map with string key. We accept a generic for the value type
since we may want to create maps with many different types of values.
*/
func CreateMap[T any]() map[string]T {
	myMap := make(map[string]T)
	return myMap
}

/*
Slices are used all the time. Especially useful when pulling collections from
a database.
*/
func WorkingWithSlices() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceOfSlice := slice[2:6]
	fmt.Println(sliceOfSlice)
}

/*
Can range over slices
*/
func ForRangeSlice() {
	slice := []int{65, 3, 41, 972, 14, 76}
	for i, x := range slice {
		fmt.Println("index:", i, "--", "value:", x)
	}
}

/*
Can range over maps
*/
func ForRangeMap() {
	myMap := make(map[string]string)
	myMap["test"] = "string1"
	myMap["new"] = "string2"
	for i, x := range myMap {
		fmt.Println("key:", i, "--", "value:", x)
	}
}
