package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LineController struct{}

func NewLineController() *LineController {
	return new(LineController)
}

func (lc *LineController) Webhook(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read request body",
		})
		return
	}
	fmt.Println(string(body))
	c.JSON(http.StatusOK, gin.H{"Status": "Hello Line :), Everything seems Ok"})
	return
}
