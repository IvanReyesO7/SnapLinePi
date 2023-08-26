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
				log.Println(event)
				log.Println(message.Text)
				if strings.Contains(strings.ToLower(message.Text), "snapshot") {
					log.Println("Initialize...")
					snapshot, err := services.TakeSnapshot()
					log.Println(snapshot)
					if err != nil {
						// TODO send failed Line msg
					}

					// Send picture to Line
					hostname := os.Getenv("HOSTNAME")
					imageUrl := hostname + *snapshot
					log.Println(imageUrl)
					message := linebot.NewImageMessage(imageUrl, imageUrl)
					// append some message to messages
					replyToken := event.ReplyToken
					messages := []linebot.SendingMessage{message}
					m, err := lc.Bot.ReplyMessage(replyToken, messages...).Do()
					log.Println(m)
					if err != nil {
						// Do something when some bad happened
						log.Println(err)
					}

					// Delete picture
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Hello Line :), Everything seems Ok"})
	return
}
