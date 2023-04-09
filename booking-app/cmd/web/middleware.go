package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

/*
This is the format for writing custom middleware. Custom middleware functions always take an http.Handler
amd return the same. http.HandlerFunc calls the callback and afterwards passes the request onwards
*/
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("You requested... %s\n", r.URL)
		next.ServeHTTP(w, r)
	})
}

/*
Create a new nosurf token. to provider CSRF protection to POST requests. Path tells us it can be
used on the entire site, Secure is disabled now since for development we don't have https.
*/
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

/*
We need a middleware to make our application aware of sessions. We can do this using the LoadAndSave
middleware that scs provides.
*/
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
