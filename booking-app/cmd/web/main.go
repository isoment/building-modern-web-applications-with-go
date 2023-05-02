package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/isoment/booking-app/internal/config"
	"github.com/isoment/booking-app/internal/handlers"
	"github.com/isoment/booking-app/internal/helpers"
	"github.com/isoment/booking-app/internal/models"
	"github.com/isoment/booking-app/internal/render"
)

const portNumber = ":8008"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting Application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// We need to define the non-primitive types we want to store in the session.
	gob.Register(models.Reservation{})

	app.InProduction = false

	// Set up logging
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Create a new session, sessions timeout after 24 hours and the cookie
	// will persist after the browser is closed. Set some other values and put
	// it in our config so we can access it anywhere.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// Create a template cache and store it in the AppConfig, we can use the
	// UseCache value to toggle use of the cache for development mode.
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.UseCache = false
	app.TemplateCache = tc

	// Create and set the handlers repository.
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Passing the application config where it is needed.
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
