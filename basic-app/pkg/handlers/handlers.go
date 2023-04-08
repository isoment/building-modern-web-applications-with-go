package handlers

import (
	"net/http"

	"github.com/isoment/basic-app/pkg/config"
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

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.html")
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.page.html")
}
