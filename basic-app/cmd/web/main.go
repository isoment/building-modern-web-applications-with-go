package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isoment/basic-app/pkg/config"
	"github.com/isoment/basic-app/pkg/handlers"
	"github.com/isoment/basic-app/pkg/render"
)

const portNumber = ":8008"

func main() {
	// Create a template cache and store it in the AppConfig
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
