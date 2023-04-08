package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// Define a map of functions that can be used in a template
var functions = template.FuncMap{}

func RenderTemplate(res http.ResponseWriter, t string) {
	_, err := RenderTemplateTest(res)
	if err != nil {
		fmt.Println("Error getting template cache")
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + t)

	err = parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func RenderTemplateTest(res http.ResponseWriter) (map[string]*template.Template, error) {
	// Define a map where the key is a string and the value is a pointer to a template
	myCache := map[string]*template.Template{}

	// Get all the template pages but not the layouts
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Loop over the page html files pages
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)

		// Create a new template for each page, attach any functions
		// we want to use
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Look for any layout files in the templates directory
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// If we found a layout parse it
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// Add the template to the map we defined above.
		myCache[name] = ts
	}

	return myCache, nil
}
