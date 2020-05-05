package models

import (
	"database/sql"

	"github.com/musale/snippets/pkg/models"
)

// UserModel struct
type UserModel struct {
	DB *sql.DB
}

// Insert a newuser
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate a user
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get a user given an ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
