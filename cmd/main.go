package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"Gold-Rate-Analyser/internal/fetcher"
	"Gold-Rate-Analyser/internal/notifier"
	"Gold-Rate-Analyser/internal/storage"
)

func main() {
	fmt.Println("[MAIN] Gold Notifier App Started ")
	fmt.Println("[MAIN] Current time:", time.Now())
	fmt.Println("[MAIN] Loading Env variables ")
	if err := godotenv.Load(); err != nil {
		log.Println("[MAIN] .env not found, using environment variables")
	}
	fmt.Println("[MAIN] Taking Market Snapshot")

	snapshot, err := fetcher.GetMarketSnapshot()
	if err != nil {
		log.Fatal(err)
	}

	err = storage.AppendSnapshot(snapshot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Snapshot appended successfully")

	message := fmt.Sprintf(
		`🏆 GOLD MARKET UPDATE

💰 Spot Gold: ₹%.2f/g
📈 MCX Gold: ₹%.2f/g
🏅 IBJA Gold: ₹%.2f/g

🌍 LBMA AM: ₹%.2f/g
🌍 LBMA PM: ₹%.2f/g

🥈 Silver: ₹%.2f/g

🕒 Updated:
%s
`,
		snapshot.Metals.Gold,
		snapshot.Metals.MCXGold,
		snapshot.Metals.IBJAGold,
		snapshot.Metals.LBMAGoldAM,
		snapshot.Metals.LBMAGoldPM,
		snapshot.Metals.Silver,
		snapshot.Timestamps.Metal,
	)

	chatIDs, err := storage.LoadChatIDs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[MAIN] Sending telegram message...")

	for _, chatID := range chatIDs {

		err := notifier.SendTelegramMessage(
			chatID,
			message,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("[MAIN] Message sent to %d\n", chatID)
	}

	fmt.Println("Telegram message sent successfully")

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

	fmt.Println("\n--- SILVER ---")
	fmt.Printf("Silver: %.2f\n", snapshot.Metals.Silver)

	fmt.Println("\n--- TIMESTAMPS ---")
	fmt.Printf("Metal Timestamp: %s\n", snapshot.Timestamps.Metal)
	fmt.Printf("Currency Timestamp: %s\n", snapshot.Timestamps.Currency)

	fmt.Println("=====================================")
}
