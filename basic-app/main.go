package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8008"

func Home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.html")
}

func About(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "about.page.html")
}

func renderTemplate(res http.ResponseWriter, t string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + t)
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
