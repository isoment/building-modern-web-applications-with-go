package main

import (
	"fmt"
	"net/http"

	"github.com/isoment/basic-app/pkg/handlers"
)

const portNumber = ":8008"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
