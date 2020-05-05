package models

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/musale/snippets/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// UserModel struct
type UserModel struct {
	DB *sql.DB
}

// Insert a newuser
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

	// Use the Exec() method to insert the user details and hashed password
	// into the users table. If this returns an error, we try to type assert
	// it to a *mysql.MySQLError object so we can check if the error number is
	// 1062 and, if it is, we also check whether or not the error relates to
	// our users_uc_email key by checking the contents of the message string.
	// If it does, we return an ErrDuplicateEmail error. Otherwise, we just
	// return the original error (or nil if everything worked).
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

// Authenticate a user
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get a user given an ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
