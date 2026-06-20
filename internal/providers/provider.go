package providers

import "Gold-Rate-Analyser/internal/models"

type Provider interface {
	Name() string
	FetchMarketSnapshot() (*models.MarketSnapshot, error)
}
