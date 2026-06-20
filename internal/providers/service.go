package providers

import (
	"fmt"
	"log"

	"Gold-Rate-Analyser/internal/models"
)

func GetMarketSnapshot() (*models.MarketSnapshot, error) {

	providers := []Provider{
		AlphaVantageProvider{},
		// YahooProvider{},
		// MetalsDevProvider{},
	}

	for _, provider := range providers {

		snapshot, err := provider.FetchMarketSnapshot()

		if err == nil {
			log.Printf(
				"[PROVIDER] Using %s",
				provider.Name(),
			)

			return snapshot, nil
		}

		log.Printf(
			"[PROVIDER] %s failed: %v",
			provider.Name(),
			err,
		)
	}

	return nil, fmt.Errorf("all providers failed")
}
