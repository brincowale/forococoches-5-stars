package main

import (
	"forococoches-5-stars/models"
	"forococoches-5-stars/parser"
	"forococoches-5-stars/utils"
	"github.com/getsentry/sentry-go"
	"strings"
)

func main() {
	configs := utils.ReadConfig()
	_ = sentry.Init(sentry.ClientOptions{
		Dsn: configs.Sentry,
	})
	utils.CreateConnectionDB(configs.DBConnection)
	for _, categoryId := range configs.Categories {
		URL := "https://www.forocoches.com/foro/forumdisplay.php?f=" + categoryId + "&daysprune=1&order=desc&sort=voteavg"
		threads := parser.Parse(URL, configs.ProxyURI)
		var validThreads []models.Thread
		for _, thread := range threads {
			if utils.IsNewThread(thread) {
				utils.InsertThread(thread)
				validThreads = append(validThreads, thread)
			}
		}
		if len(validThreads) > 0 {
			message := strings.TrimSpace(utils.CreateMessage(validThreads))
			utils.SendTelegramMessage(message, configs)
		}
	}
}
