package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasskadd/Event-management/auth-service/models"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register handles user registration
func Register(c *gin.Context) {
	var request RegistrationRequest

	// Bind the JSON input to the request struct
	if err := c.BindJSON(&request); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get the DB connection from the context (passed from middleware)
	var db *sql.DB = c.MustGet("db").(*sql.DB)
	fmt.Println("Database connection established.")

	// Try to register the user
	success, errors := models.RegisterUser(db, request.Username, request.Email, request.Password)

	if !success {
		fmt.Println("User registration failed:", errors)
		var errorMessages []string
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var db *sql.DB = c.MustGet("db").(*sql.DB)

	//Validate Login
	success, user, errors := models.ValidateLogin(db, request.Email, request.Password)

	if !success {
		fmt.Println("User login failed:", errors)
		var errorMessages []string
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	// Generate JWT token
	token, err := models.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Create the cookie with the token
	models.CreateCookie(c, token)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "token": token})
}

func Logout(c *gin.Context) {

}
