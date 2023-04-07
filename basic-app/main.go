package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8008"

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This is the homepage")
}

func About(res http.ResponseWriter, req *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(res, fmt.Sprintf("This is the about page 2 + 2 is %d", sum))
}

func Divide(res http.ResponseWriter, req *http.Request) {
	x := float32(100.00)
	y := float32(0.0)
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(res, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(res, fmt.Sprintf("%f / %f = %f", x, y, f))
}

func addValues(x, y int) int {
	sum := x + y
	return sum
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
