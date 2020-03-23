package utils

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBConnection    string
	Categories      []string
	TelegramApiKey  string
	TelegramChannel string
	Sentry          string
	ProxyURI        string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(time.Second * 5)
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config := Config{
		DBConnection:    viper.GetString("database.connection"),
		Categories:      viper.GetStringSlice("fc_categories"),
		TelegramApiKey:  viper.GetString("telegram.api_key"),
		TelegramChannel: viper.GetString("telegram.channel"),
		Sentry:          viper.GetString("sentry"),
		ProxyURI:        viper.GetString("proxy"),
	}
	return config
}
