package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord when no matching record is found
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials when credentials are wrong
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail when email already exists
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet is a single snip
type Snippet struct {
	ID               int
	Title, Content   string
	Created, Expires time.Time
}

// User is a single user
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
