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
Create a new nosurf token. Path tells us it can be sued on the entire site, Secure is disabled now
since for development we don't have TLS
*/
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
