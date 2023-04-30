package render

import (
	"net/http"
	"testing"

	"github.com/isoment/booking-app/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// Put something in the session
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc

	// Create a request with session data
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	// We should be able to render a template that exists
	err = RenderTemplate(&ww, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("Error rendering template")
	}

	// We should not be able to render a template that does not exist
	err = RenderTemplate(&ww, r, "fake.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("Rendered template that does not exist")
	}

	// Testing rendering templates with the cache enabled
	app.UseCache = true
	err = RenderTemplate(&ww, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("Error rendering template")
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

/*
We need to create a request and get the request context. We the use the
context to get the session data from the store. The context is used for
sharing values that are unique to a specific single request
*/
func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	// Load session token from request header
	sessionToken := r.Header.Get("X-Session")
	// Load the session object into the context
	ctx, _ = session.Load(ctx, sessionToken)
	r = r.WithContext(ctx)
	return r, nil
}
