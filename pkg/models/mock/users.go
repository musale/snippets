package mock

import (
	"time"

	"github.com/musale/snippets/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Kunta",
	Email:   "kunta@snippets.com",
	Created: time.Now(),
}

// UserModel -
type UserModel struct{}

// Insert into user
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate a user
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "kunta@snippets.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get a User
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
