package main

import (
	"path/filepath"
	"text/template"

	"github.com/allensuvorov/snippetbox.git/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(".ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	return cache, nil
}
