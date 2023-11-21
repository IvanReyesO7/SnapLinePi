package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server-test/services"
	"strings"
	"time"

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
				switch {
				case strings.Contains(strings.ToLower(message.Text), "snapshot"):
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}

					flexMessage, err := snapshot()
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
					replyToken := event.ReplyToken
					_, err = lc.Bot.ReplyMessage(replyToken, flexMessage).Do()
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Hello Line :), Everything seems Ok"})
	return
}

func snapshot() (*linebot.FlexMessage, error) {
	log.Println("Initialize...")
	snapshot, err := services.TakeSnapshot()
	if err != nil {
		return nil, err
	}

	hostname := os.Getenv("HOSTNAME")
	imageUrl := hostname + *snapshot

	currentTime := time.Now().Add(9 * time.Hour)
	formattedDate := currentTime.Format("2006/01/02")
	formattedTime := currentTime.Format("15:04:05")

	text := fmt.Sprintf("%s at %s", formattedDate, formattedTime)

	flexMessage := services.BuildFlexMessage(imageUrl, text)
	return flexMessage, nil
}
