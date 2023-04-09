package handlers

import (
	"net/http"

	"github.com/isoment/basic-app/pkg/config"
	"github.com/isoment/basic-app/pkg/models"
	"github.com/isoment/basic-app/pkg/render"
)

// The repository used by the handlers
var Repo *Repository

// The Repository type, contains a pointer to the applications config
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository, return a pointer to the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Create some string data to pass into the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
