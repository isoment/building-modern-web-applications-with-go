package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Define a map of functions that can be used in a template
var functions = template.FuncMap{}

/*
Render a given template.
*/
func RenderTemplate(res http.ResponseWriter, template string) {
	// Get the template cache from the app config
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Find the requested template in the cache
	t, ok := tc[template]
	if !ok {
		log.Fatal(err)
	}

	// Create a new buffer in memory for manipulating byte data. We can call execute on
	// the template passing in the buffer.
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// We write the buffer data to the http.ResponseWriter
	_, err = buf.WriteTo(res)
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
