package notifier

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendTelegramMessage(
	chatID int64,
	message string,
) error {

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		token,
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