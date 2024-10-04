package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yasskadd/Event-management/auth-service/handlers"
)

// SetupRoutes initializes the routes for the auth service
func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
