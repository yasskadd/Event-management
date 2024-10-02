package models

import (
	"database/sql"
	"regexp"
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func RegisterUser(db *sql.DB, username string, email string, password string) error {
	// Validate the username and password with Regex

	// Check if email is already taken

	//Check if username is already taken

	// If everythings good, hash the password and add user to DB
	return nil
}

func IsUsernameValid(username string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`)
	return reg.MatchString(username)
}

func isEmailValid(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(email)
}

func IsUsernameTaken()

func IsEmailTaken()
