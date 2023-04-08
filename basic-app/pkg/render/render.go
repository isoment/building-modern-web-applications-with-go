package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(res http.ResponseWriter, t string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + t)
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
