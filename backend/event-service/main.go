package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {

	// Database connection
	host := "postgres"
	port := 5432
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	// Ping the database to ensure a successful connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v\n", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Pass db to Gin context
	//router.Use(middleware.DBMiddleware(db))

	// Setup routes for event-service
	//routes.SetupRoutes(router)

	// Log message indicating server is starting
	log.Println("Starting event service on http://localhost:8081")

	// Start the server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
