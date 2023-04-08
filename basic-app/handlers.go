package main

import (
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.html")
}

func About(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "about.page.html")
}
