package fetcher

func GetMarketSnapshot() (*MarketSnapshot, error) {
	return FetchMarketSnapshot()
}