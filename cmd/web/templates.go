package main

import "github.com/musale/snippets/pkg/models"

// templateData is a holding for all dynamic data
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
