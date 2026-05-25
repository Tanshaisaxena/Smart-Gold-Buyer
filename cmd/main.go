package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"Gold-Rate-Analyser/internal/fetcher"
	"Gold-Rate-Analyser/internal/storage"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	snapshot, err := fetcher.GetMarketSnapshot()
	if err != nil {
		log.Fatal(err)
	}
	err = storage.AppendSnapshot(snapshot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Snapshot appended successfully")

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
