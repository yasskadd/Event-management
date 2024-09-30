package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register handles user registration
func Register(c *gin.Context) {
	// Logic for user registration
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

// Login handles user login
func Login(c *gin.Context) {
	// Logic for user login
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!"})
}
