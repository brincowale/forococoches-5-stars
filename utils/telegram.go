package utils

import (
	"forococoches-5-stars/models"
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
	"net/http"
	"time"
)

var request *gorequest.SuperAgent

func init() {
	request = gorequest.New().
		Set("Content-Type", "application/json").
		Timeout(30*time.Second).
		Retry(3, 5*time.Second, http.StatusInternalServerError)
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendTelegramMessage(message string, config Config) bool {
	var URL = "https://api.telegram.org/bot" + config.TelegramApiKey + "/sendMessage"
	data := Message{
		ChatId: config.TelegramChannel,
		Text:   message,
	}
	_, body, errList := request.Post(URL).Send(data).End()
	if errList != nil {
		for _, err := range errList {
			sentry.CaptureException(err)
		}
	}
	if gjson.Get(body, "ok").Bool() {
		return true
	}
	sentry.CaptureException(errors.New("Cannot send to Telegram"))
	return false
}

func CreateMessage(threads []models.Thread) string {
	var message string
	for _, thread := range threads {
		returnLine := "\n"
		message += thread.Title + returnLine + thread.URL + returnLine + returnLine
	}
	return message
}
