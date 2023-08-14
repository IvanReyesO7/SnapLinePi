package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running!"})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Server is up and running!"})
	})
	r.POST("/line_webhook", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "Ok"})
	})
	r.Run()
}
