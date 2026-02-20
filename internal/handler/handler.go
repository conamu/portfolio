package handler

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/conamu/portfolio/internal/config"
)

// TemplateData is the common envelope passed to every template.
type TemplateData struct {
	Theme    config.Theme
	Page     string // "home" | "about"
	Projects []config.Project
	About    config.About
}

// Handler holds shared dependencies for all HTTP handlers.
type Handler struct {
	templates *template.Template
	static    http.Handler
	theme     config.Theme
	projects  []config.Project
	about     config.About
}

func New(fsys fs.FS, theme config.Theme, projects []config.Project, about config.About) *Handler {
	funcMap := template.FuncMap{
		"nl2br": func(s string) template.HTML {
			escaped := template.HTMLEscapeString(s)
			return template.HTML(strings.ReplaceAll(escaped, "\n", "<br>"))
		},
	}

	tmplFS, err := fs.Sub(fsys, "web/templates")
	if err != nil {
		panic("assets: cannot sub into web/templates: " + err.Error())
	}
	tmpl := template.Must(
		template.New("").Funcs(funcMap).ParseFS(tmplFS, "*.html"),
	)

	staticFS, err := fs.Sub(fsys, "web/static")
	if err != nil {
		panic("assets: cannot sub into web/static: " + err.Error())
	}

	return &Handler{
		templates: tmpl,
		static:    http.FileServerFS(staticFS),
		theme:     theme,
		projects:  projects,
		about:     about,
	}
}

// Static returns the embedded static file handler (strip /static/ before using).
func (h *Handler) Static() http.Handler { return h.static }

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.render(w, "home.html", TemplateData{
		Theme:    h.theme,
		Page:     "home",
		Projects: h.projects,
	})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.render(w, "about.html", TemplateData{
		Theme: h.theme,
		Page:  "about",
		About: h.about,
	})
}

func (h *Handler) render(w http.ResponseWriter, name string, data TemplateData) {
	if err := h.templates.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("template error (%s): %v", name, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
