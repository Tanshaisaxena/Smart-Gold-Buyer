package configs

import (
	"fmt"
	"os"
)

type Config struct {
	Apikey string `json:"METAL_API_KEY"`
	Telegram_Bot_Token string `json:"TELEGRAM_BOT_TOKEN"`
}

func Configloader() (Config, error) {
	var config Config
	ApiKeyvar := os.Getenv("METALS_API_KEY")
	if ApiKeyvar != "" {
		config.Apikey = ApiKeyvar
	} else {
		return config, fmt.Errorf("[ConfigLoader] Apikey loading failed")
	}

	Telegram_Bot_Token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if Telegram_Bot_Token!=""{
		config.Telegram_Bot_Token=Telegram_Bot_Token
	} else {
		return config, fmt.Errorf("[ConfigLoader] Telegram Bot token loading failed")
	}

	return config, nil
}
