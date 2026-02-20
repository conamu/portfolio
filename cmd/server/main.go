package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/conamu/portfolio"
	"github.com/conamu/portfolio/internal/config"
	"github.com/conamu/portfolio/internal/handler"
)

func main() {
	// The embed FS roots at the repo root, so data files live at "data/...".
	dataFS, err := fs.Sub(portfolio.FS, "data")
	if err != nil {
		log.Fatalf("assets sub data: %v", err)
	}

	theme, err := config.LoadTheme(dataFS, "theme.json")
	if err != nil {
		log.Fatalf("load theme: %v", err)
	}

	projects, err := config.LoadProjects(dataFS, "projects.json")
	if err != nil {
		log.Fatalf("load projects: %v", err)
	}

	about, err := config.LoadAbout(dataFS, "about.json")
	if err != nil {
		log.Fatalf("load about: %v", err)
	}

	h := handler.New(portfolio.FS, theme, projects, about)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /about", h.About)
	mux.HandleFunc("GET /cv", h.CV)
	mux.Handle("GET /static/", http.StripPrefix("/static/", h.Static()))

	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
