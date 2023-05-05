package handlers

import (
	"net/http"

	"github.com/kabaf81/BuildAWebApplication/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl.html")
}

func SiteMap(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "site.page.tmpl.html")
}
