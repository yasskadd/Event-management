package models

import (
	"database/sql"
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
	return nil
}
