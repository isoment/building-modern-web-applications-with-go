package main

import (
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

func addValues(x, y int) int {
	sum := x + y
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
