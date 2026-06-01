package notifier

import (
	"Gold-Rate-Analyser/internal/configs"
	"fmt"
	"net/http"
	"net/url"
)

func SendTelegramMessage(
	config *configs.Config,
	chatID int64,
	message string,
) error {

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		config.Telegram_Bot_Token,
	)

	data := url.Values{}

	data.Set("chat_id", fmt.Sprintf("%d", chatID))
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf(
			"telegram API failed with status %s",
			resp.Status,
		)
	}

	return nil
}
