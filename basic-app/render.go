package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(res http.ResponseWriter, t string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + t)
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
