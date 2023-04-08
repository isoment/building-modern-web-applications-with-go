package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/isoment/basic-app/pkg/config"
)

// Define a map of functions that can be used in a template
var functions = template.FuncMap{}

var app *config.AppConfig

// Sets the config for the template
func NewTemplates(a *config.AppConfig) {
	app = a
}

/*
Render a given template.
*/
func RenderTemplate(res http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	// If the application is in development we may not want to use the template cache,
	// we include this check to see if it is disabled in the application config.
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Find the requested template in the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Create a new buffer in memory for manipulating byte data. We can call execute on
	// the template passing in the buffer.
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// We write the buffer data to the http.ResponseWriter
	_, err := buf.WriteTo(res)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

/*
We want to parse all of the templates including the layouts and store them in a map, this
will be our template cache.
*/
func CreateTemplateCache() (map[string]*template.Template, error) {
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

		myCache[name] = ts
	}

	return myCache, nil
}
