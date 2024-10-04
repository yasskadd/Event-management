package models

import (
	"os"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
