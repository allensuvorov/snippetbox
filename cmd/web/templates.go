package main

import (
	"text/template"

	"github.com/allensuvorov/snippetbox.git/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	return cache, nil
}
