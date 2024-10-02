package models

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/yasskadd/Event-management/auth-service/utils"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func ValidateRegistration(db *sql.DB, username string, email string, password string) []error {
	var errorsList []error
	// Validate the username and password with Regex
	if !IsUsernameValid(username) {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeInvalidUsername, utils.ErrInvalidUsername))
	}
	if !IsEmailValid(email) {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeInvalidEmail, utils.ErrInvalidEmail))
	}

	// Check if password is valid
	if !IsPasswordValid(password) {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodePasswordTooWeak, utils.ErrPasswordTooWeak))
	}

	// Check if username is already taken
	if taken, _ := IsUsernameTaken(db, username); taken {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeUsernameAlreadyTaken, utils.ErrUsernameAlreadyTaken))
	}

	// Check if email is already taken
	if taken, _ := IsEmailTaken(db, email); taken {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeEmailAlreadyTaken, utils.ErrEmailAlreadyTaken))
	}

	// If everythings good, hash the password and add user to DB
	return errorsList
}

func IsUsernameValid(username string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`)
	return reg.MatchString(username)
}

func IsEmailValid(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(email)
}

func IsPasswordValid(password string) bool {
	reg := regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*(),.?":{}|<>])[A-Za-z\d!@#$%^&*(),.?":{}|<>]{8,}$`)
	return reg.MatchString(password)
}

func IsUsernameTaken(db *sql.DB, username string) (bool, error) {
	var exists bool
	const query string = "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func IsEmailTaken(db *sql.DB, email string) (bool, error) {
	var exists bool
	const query string = "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}
