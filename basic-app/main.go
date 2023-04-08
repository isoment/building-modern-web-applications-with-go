package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8008"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
