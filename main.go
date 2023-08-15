package main

import (
	"fmt"
	"log"
	"net/http"
	"server-test/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	lineController, err := controllers.NewLineController()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not load Line controller, err: %e", err))
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running!"})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running!"})
	})
	r.POST("/line_webhook", lineController.Webhook)
	r.Run()
}
