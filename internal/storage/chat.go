package storage

import (
	"encoding/json"
	"os"
)

const chatFile = "data/chat_ids.json"

func LoadChatIDs() ([]int64, error) {

	var chatIDs []int64

	file, err := os.ReadFile(chatFile)
	if err != nil {
		if os.IsNotExist(err) {
			return chatIDs, nil
		}
		return chatIDs, err
	}

	err = json.Unmarshal(file, &chatIDs)
	if err != nil {
		return chatIDs, err
	}

	return chatIDs, nil
}

func SaveChatID(chatID int64) error {

	chatIDs, err := LoadChatIDs()
	if err != nil {
		return err
	}

	// Prevent duplicates
	for _, id := range chatIDs {
		if id == chatID {
			return nil
		}
	}

	chatIDs = append(chatIDs, chatID)

	data, err := json.MarshalIndent(chatIDs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(chatFile, data, 0644)
}
