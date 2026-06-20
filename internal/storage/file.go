package storage

import (
	"encoding/json"
	"os"

	"Gold-Rate-Analyser/internal/models"
)

const outputFile = "data/market_snapshots.jsonl"

func AppendSnapshot(snapshot *models.MarketSnapshot) error {

	file, err := os.OpenFile(
		outputFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		return err
	}

	return nil
}