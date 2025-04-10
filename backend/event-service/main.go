package main

import (
	"event-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up the routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	log.Println("Starting event-service on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Println("Error")
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server succesfully started")
}
