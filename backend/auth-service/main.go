package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //PostgreSQL Driver
	"github.com/yasskadd/Event-management/auth-service/middleware"
	"github.com/yasskadd/Event-management/auth-service/routes"
)

func main() {

	host := "postgres"
	port := 5432
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the DB
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v\n", err)
	}

	router := gin.Default()

	//Pass db to gin context
	router.Use(middleware.DBMiddleware(db))

	// Setup routes
	routes.SetupRoutes(router)

	// Log message indicating server is starting
	log.Println("Starting server on http://localhost:8080")

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
