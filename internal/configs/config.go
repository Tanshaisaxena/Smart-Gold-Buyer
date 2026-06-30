package configs

import (
	"fmt"
	"os"
)

type Config struct {
	ALPHA_VANTAGE_API_KEY string `json:"ALPHA_VANTAGE_API_KEY"`
	Telegram_Bot_Token    string `json:"TELEGRAM_BOT_TOKEN"`
}

func getEnvOrFallback(primary string, fallbacks ...string) string {
	if value := os.Getenv(primary); value != "" {
		return value
	}

	for _, key := range fallbacks {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return ""
}

func Configloader() (Config, error) {
	var config Config
	apiKey := getEnvOrFallback("ALPHA_VANTAGE_API_KEY", "METALS_API_KEY")
	if apiKey != "" {
		config.ALPHA_VANTAGE_API_KEY = apiKey
	} else {
		return config, fmt.Errorf("[ConfigLoader] Apikey loading failed")
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken != "" {
		config.Telegram_Bot_Token = telegramBotToken
	} else {
		return config, fmt.Errorf("[ConfigLoader] Telegram Bot token loading failed")
	}

	return config, nil
}
