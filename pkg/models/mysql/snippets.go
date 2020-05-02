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
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
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
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM
	snippets WHERE expires > UTC_TIMESTAMP() AND id=?`
	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

// Latest will return the 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
