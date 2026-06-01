package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"Gold-Rate-Analyser/internal/configs"
	"Gold-Rate-Analyser/internal/constants"
	"Gold-Rate-Analyser/internal/fetcher"
	"Gold-Rate-Analyser/internal/message"
	"Gold-Rate-Analyser/internal/notifier"
	"Gold-Rate-Analyser/internal/storage"
)

func main() {
	fmt.Println("[MAIN] Gold Notifier App Started ")
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

	slog.SetDefault(logger)

	logger.Debug("Test Log")

	logger.Info("System", "Current time", time.Now().UTC())
	timeinindia, err := time.LoadLocation(constants.Location)
	if err != nil {
		logger.Error("ERROR while loading location", "Err", err)
	}
	logger.Info("Time Print", "Current time", time.Now().In(timeinindia))

	logger.Info("Loading Env variables ")
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}
	Config, err := configs.Configloader()
	if err != nil {
		logger.Error("Error loading config: ", "Err", err)
	}
	fmt.Println("Taking Market Snapshot")

	snapshot, err := fetcher.GetMarketSnapshot(&Config)
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

	logger.Info("Message Sent to subscribers",
		"Telegram Message Sent to %d subscribers\n",
		len(chatIDs),
	)

	fmt.Println("========== MARKET SNAPSHOT ==========")

	fmt.Printf("Currency: %s\n", snapshot.Currency)
	fmt.Printf("Unit: %s\n", snapshot.Unit)

	fmt.Println("\n--- GOLD ---")
	fmt.Printf("Spot Gold: %.2f\n", snapshot.Metals.Gold)

	fmt.Println("\n--- MCX ---")
	fmt.Printf("MCX Gold: %.2f\n", snapshot.Metals.MCXGold)
	fmt.Printf("MCX Gold AM: %.2f\n", snapshot.Metals.MCXGoldAM)
	fmt.Printf("MCX Gold PM: %.2f\n", snapshot.Metals.MCXGoldPM)

	fmt.Println("\n--- IBJA ---")
	fmt.Printf("IBJA Gold: %.2f\n", snapshot.Metals.IBJAGold)

	fmt.Println("\n--- LBMA ---")
	fmt.Printf("LBMA Gold AM: %.2f\n", snapshot.Metals.LBMAGoldAM)
	fmt.Printf("LBMA Gold PM: %.2f\n", snapshot.Metals.LBMAGoldPM)

	fmt.Println("\n--- Related Metals ---")
	fmt.Printf("Silver: %.2f\n", snapshot.Metals.Silver)
	fmt.Printf("Platinium: %.2f\n", snapshot.Metals.Platinum)
	fmt.Printf("Palladium: %.2f\n", snapshot.Metals.Palladium)

	fmt.Println("\n--- TIMESTAMPS ---")
	fmt.Printf("Metal Timestamp: %s\n", snapshot.Timestamps.Metal)
	fmt.Printf("Currency Timestamp: %s\n", snapshot.Timestamps.Currency)

	fmt.Println("=====================================")
}
