package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yasskadd/Event-management/auth-service/handlers"
)

type RegistrationRequest struct {
	Username string `json:"username"` // Tags can be used to specify how to marshal/unmarshal JSON
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SetupRoutes initializes the routes for the auth service
func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register) // Placeholder for registration handler
	router.POST("/login", handlers.Login)       // Placeholder for login handler
}
