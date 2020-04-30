package models

import (
	"database/sql"

	"github.com/musale/snippets/pkg/models"
)

// SnippetModel wraps a DB connection pool for Snippets
type SnippetModel struct {
	DB *sql.DB
}

// Insert makes a new Snippet addition in the DB
func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get returns a snippet if found
func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recent snippets
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
