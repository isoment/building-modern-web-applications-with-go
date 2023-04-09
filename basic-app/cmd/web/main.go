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
	var app config.AppConfig

	// Create a template cache and store it in the AppConfig, we can use the
	// UseCache value to toggle use of the cache for development mode.
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.UseCache = false
	app.TemplateCache = tc

	// Create and set the handlers repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting Application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
