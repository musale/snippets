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
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

// Get returns a snippet if found
func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recent snippets
func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
