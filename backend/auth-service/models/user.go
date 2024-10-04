package models

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"github.com/yasskadd/Event-management/auth-service/dao"
	"github.com/yasskadd/Event-management/auth-service/utils"
)

func ValidateLogin(db *sql.DB, email string, password string) (bool, *dao.User, []error) {
	var errorsList []error

	if !IsEmailValid(email) {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeInvalidEmail, utils.ErrInvalidEmail))
		return false, nil, errorsList
	}

	user, err := dao.GetUserByEmail(db, email)
	if err != nil {
		errorsList = append(errorsList, fmt.Errorf("failed to retrieve user: %v", err))
		return false, nil, errorsList
	}
	if user == nil {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeUserNotFound, utils.ErrUserNotFound))
		return false, nil, errorsList
	}
	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeInvalidPassword, utils.ErrInvalidPassword))
		return false, nil, errorsList
	}

	return true, user, nil
}

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
		return false, []error{utils.NewAuthentificationError(utils.ErrCodeHashingError, utils.ErrHashingError)} // Return hashing error
	}

	const query string = "INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, username, email, hashedPassword, time.Now())
	if err != nil {
		fmt.Println("Error inserting user into DB:", err)
		return false, []error{utils.NewAuthentificationError(utils.ErrCodeDatabaseError, utils.ErrDatabaseError)}
	}

	return true, nil
}

func ValidateRegistration(db *sql.DB, username string, email string, password string) []error {
	var errorsList []error
	if !IsUsernameValid(username) {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeInvalidUsername, utils.ErrInvalidUsername))
	}
	if !IsEmailValid(email) {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeInvalidEmail, utils.ErrInvalidEmail))
	}
	if !IsPasswordValid(password) {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodePasswordTooWeak, utils.ErrPasswordTooWeak))
	}
	if taken, _ := IsUsernameTaken(db, username); taken {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeUsernameAlreadyTaken, utils.ErrUsernameAlreadyTaken))
	}
	if taken, _ := IsEmailTaken(db, email); taken {
		errorsList = append(errorsList, utils.NewAuthentificationError(utils.ErrCodeEmailAlreadyTaken, utils.ErrEmailAlreadyTaken))
	}
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
