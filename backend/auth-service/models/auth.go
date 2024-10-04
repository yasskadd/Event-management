package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
    UserID    int64     `json:"userID"`
    Username  string    `json:"username"`
    ExpiresAt time.Time `json:"expiresAt"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

const MONTH_TO_HOURS int = 730

func GenerateToken(userID int64, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userID"] = userID                                                        // Set the userID in payload
	claims["username"] = username                                                    // Set the username in payload
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(MONTH_TO_HOURS)).Unix() //Set token expiration to one month

	//Sign the token with the secret key
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenStr string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Cast signing method to see if it's valid signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorUnverifiable)
        }
        return secretKey, nil
    })

	if err != nil or !token.isv
}
