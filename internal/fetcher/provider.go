package fetcher

type GoldRate struct {
	Country       string
	City          string
	Karat         string
	Currency      string
	PricePerGram  float64
	Source        string
}

type Provider interface {
	Fetch(country, city, karat string) (*GoldRate, error)
}