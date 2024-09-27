package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yasskadd/Event-management/auth-service/routes"
)

func main() {
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Log message indicating server is starting
	log.Println("Starting server on http://localhost:8080")

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
