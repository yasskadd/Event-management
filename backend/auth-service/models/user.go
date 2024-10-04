package models

import (
	"database/sql"
	"fmt"
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

func ValidateLogin(db *sql.DB, email string) (bool, []error) {
	//Email must be valid, email must exist in DB
	var errorsList []error
	if !IsEmailValid(email) {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeInvalidEmail, utils.ErrInvalidEmail))
	}
	if taken, _ := IsEmailTaken(db, email); taken {
		errorsList = append(errorsList, utils.NewRegistrationError(utils.ErrCodeEmailAlreadyTaken, utils.ErrEmailAlreadyTaken))
	}
	if len(errorsList) == 0 {
		return true, nil
	}
	return false, errorsList
}

// Register users if no errors, return (success, errors)
func RegisterUser(db *sql.DB, username string, email string, password string) (bool, []error) {
	// Validate registration input
	errors := ValidateRegistration(db, username, email, password)
	if len(errors) > 0 {
		fmt.Println("Validation errors found:", errors)
		return false, errors
	}

	// Hash the password and add user to the DB
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return false, []error{utils.NewRegistrationError(utils.ErrCodeHashingError, utils.ErrHashingError)} // Return hashing error
	}

	const query string = "INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, username, email, hashedPassword, time.Now())
	if err != nil {
		fmt.Println("Error inserting user into DB:", err)
		return false, []error{utils.NewRegistrationError(utils.ErrCodeDatabaseError, utils.ErrDatabaseError)} // Return database error
	}

	fmt.Println("User registered successfully!") // Success message
	return true, nil                             // Registration successful
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
	if len(password) < 8 {
		return false
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasLower && hasUpper && hasDigit && hasSpecial
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
