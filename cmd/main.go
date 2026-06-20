package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"Gold-Rate-Analyser/internal/configs"
	"Gold-Rate-Analyser/internal/message"
	"Gold-Rate-Analyser/internal/notifier"
	"Gold-Rate-Analyser/internal/providers"
	"Gold-Rate-Analyser/internal/storage"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("[MAIN] Gold Notifier App Started ")

	err := godotenv.Load()
	if err != nil {
		logger.Warn(".env file not found, using system environment variables")
	}
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(
					slog.TimeKey,
					a.Value.Time().Format("2006-01-02 15:04:05"),
				)
			}
			return a
		},
	}
	var logger *slog.Logger
	if strings.ToUpper(os.Getenv("ENV")) == "PROD" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	}

	snapshot, err := providers.GetMarketSnapshot()
	if err != nil {
		log.Fatal(err)
	}

	err = storage.AppendSnapshot(snapshot)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Snapshot appended successfully")

	Message := message.CreateMessage(snapshot)

	chatIDs, err := storage.LoadChatIDs()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("[MAIN] Sending telegram message...")

	if len(chatIDs) == 0 {
		logger.Info(" No subscribers found")
		return
	}
	Config, err := configs.Configloader()
	if err != nil {
		log.Fatal(err)
	}

	for _, chatID := range chatIDs {

		err := notifier.SendTelegramMessage(
			&Config,
			chatID,
			Message,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("[MAIN] Message sent to %d\n", chatID)
	}

	logger.Info(
		"Messages sent",
		"subscriber_count",
		len(chatIDs),
	)

	fmt.Println("========== MARKET SNAPSHOT ==========")

	fmt.Printf("Currency: %s\n", snapshot.Currency)
	fmt.Printf("Unit: %s\n", snapshot.Unit)

	fmt.Println("\n--- GOLD ---")
	fmt.Printf("Spot Gold: %.2f\n", snapshot.Metals.Gold)

	fmt.Println("\n--- Related Metals ---")
	fmt.Printf("Silver: %.2f\n", snapshot.Metals.Silver)
	fmt.Printf("Platinium: %.2f\n", snapshot.Metals.Platinum)
	fmt.Printf("Palladium: %.2f\n", snapshot.Metals.Palladium)

	fmt.Println("\n--- TIMESTAMPS ---")
	fmt.Printf("Metal Timestamp: %s\n", snapshot.Timestamps.Metal)

	fmt.Println("=====================================")
}
