package notifier

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendTelegramMessage(message string) error {

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		token,
	)

	data := url.Values{}

	data.Set("chat_id", chatID)
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