package main

import "github.com/allensuvorov/snippetbox.git/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
