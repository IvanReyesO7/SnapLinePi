package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LineController struct{}

func NewLineController() *LineController {
	return new(LineController)
}

func (lc *LineController) TestWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "Hello Line :), Everything seems Ok"})
	return
}
