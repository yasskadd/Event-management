package models

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TokenPayload struct {
	UserID    int64     `json:"userID"`
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"expiresAt"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

const MONTH_TO_HOURS int = 730

func GenerateToken(userID int64, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
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

func ValidateToken(tokenStr string) (*TokenPayload, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Cast signing method to see if it's valid signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorUnverifiable)
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	//Extract claims and verify if each claim is asserted to the correct data type (security measure)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.NewValidationError("Invalid claims", jwt.ValidationErrorUnverifiable)
	}

	userID, ok_id := claims["userID"].(float64)

	username, ok_username := claims["username"].(string)

	exp, ok_exp := claims["exp"].(float64)

	if !ok_id || !ok_username || !ok_exp {
		return nil, jwt.NewValidationError("Invalid expiration time", jwt.ValidationErrorUnverifiable)
	}

	return &TokenPayload{
		UserID:    int64(userID),
		Username:  username,
		ExpiresAt: time.Unix(int64(exp), 0),
	}, nil
}

func CreateCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Expires:  time.Now().AddDate(0, 1, 0), // One month
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Didn't implement HTTPS yet
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)
}
