package utils

type RegistrationError struct {
	Code    int
	Message string
}

// Implements Error interface
func (e *RegistrationError) Error() string {
	return e.Message
}

// Error codes
const (
	ErrCodeInvalidUsername      = iota // 0
	ErrCodeInvalidEmail                // 1
	ErrCodePasswordTooWeak             // 2
	ErrCodeEmailAlreadyTaken           // 3
	ErrCodeUsernameAlreadyTaken        // 4
	ErrCodeDatabaseError               // 5
	ErrCodeHashingError                // 6
	ErrCodeUserNotFound                // 7
	ErrCodeInvalidPassword             // 8
)

const (
	ErrInvalidUsername      = "username must be between 5 to 20 characters and can only contain letters and numbers"
	ErrInvalidEmail         = "invalid email format"
	ErrPasswordTooWeak      = "password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	ErrEmailAlreadyTaken    = "email is already taken"
	ErrUsernameAlreadyTaken = "username is already taken"
	ErrDatabaseError        = "an error occurred while interacting with the database"
	ErrHashingError         = "an error occurred while hashing the password"
	ErrUserNotFound         = "user not found"
	ErrInvalidPassword      = "invalid password"
)

// NewRegistrationError creates a new RegistrationError.
func NewRegistrationError(code int, message string) error {
	return &RegistrationError{Code: code, Message: message}
}
