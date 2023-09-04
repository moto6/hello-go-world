package handlers

import (
	"hello-world/pkg/config"
	"hello-world/pkg/models"
	"hello-world/pkg/render"
	"net/http"
)

// Repo the repositroy used by handers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo 리포
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the Repos for the handlers
func NewHandler(r *Repository) {
	Repo = r

}

// Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request, td *models.TemplateData) {
	render.RenderTemplate(w, "home.page.tmpl", td)
}

// About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request, td *models.TemplateData) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello world!"

	render.RenderTemplate(w, "about.page.tmpl", td)
}
