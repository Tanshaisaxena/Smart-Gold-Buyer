package fetcher

import "Gold-Rate-Analyser/internal/configs"

func GetMarketSnapshot(Config *configs.Config) (*MarketSnapshot, error) {
	return FetchMarketSnapshot(Config)
}