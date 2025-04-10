package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yasskadd/Event-management/auth-service/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.DELETE("logout", handlers.Logout)
}
