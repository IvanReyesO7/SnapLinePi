package main

import (
	"fmt"
	"log"
	"net/http"
	"server-test/controllers"
	"server-test/tasks"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// Run background task
	go tasks.CleanTempDir()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authentication"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	lineController, err := controllers.NewLineController()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not load Line controller, err: %e", err))
	}

	router.Static("/tmp", "./tmp")

	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running!"}) })
	router.POST("/line_webhook", lineController.Webhook)

	// r.Use(middleware.VerifyAuth)
	router.GET("/snapshot")
	router.Run()
}
