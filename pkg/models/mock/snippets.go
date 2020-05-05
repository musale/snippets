package mock

import (
	"time"

	"github.com/musale/snippets/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Test snippet",
	Content: "An old test snippet",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel -
type SnippetModel struct{}

// Insert into snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get a snippet
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest x snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
