package handlers

import (
	"net/http"

	"github.com/isoment/basic-app/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.html")
}

func About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.page.html")
}
