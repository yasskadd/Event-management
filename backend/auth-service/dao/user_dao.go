package dao

import (
	"database/sql"
)

type User struct {
	ID       int64
	Username string
	Password string
}

// GetUser retrieves a user by their ID.
func GetUserById(db *sql.DB, userID int64) (*User, error) {
	var user User
	query := "SELECT id, username, password FROM users WHERE id = $1"

	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	query := "SELECT user_id, username, password FROM users WHERE email = $1"

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}

// UserExists checks if a user exists by their username.
func UserExists(db *sql.DB, username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"

	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
