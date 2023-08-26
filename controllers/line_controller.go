package controllers

import (
	"log"
	"net/http"
	"os"
	"server-test/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineController struct {
	Bot *linebot.Client
}

func NewLineController() (*LineController, error) {

	bot, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Println("Could not valid channel credentials")
		return nil, err
	}
	lineController := LineController{Bot: bot}
	log.Println(bot)
	log.Println(lineController)
	return &lineController, nil
}

func (lc *LineController) Webhook(c *gin.Context) {
	events, err := lc.Bot.ParseRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Server internal error",
			"error":  err,
		})
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Println(message.Text)
				if strings.Contains(strings.ToLower(message.Text), "snapshot") {
					log.Println("Initialize...")
					services.TakeSnapshot()
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Hello Line :), Everything seems Ok"})
	return
}
