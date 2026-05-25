package fetcher

import (
	"fmt"
	"strings"
)

type IndiaProvider struct{}

func (i IndiaProvider) Fetch(
	country string,
	city string,
	karat string,
) (*GoldRate, error) {

	city = strings.ToLower(city)

	mockRates := map[string]float64{
		"mumbai":   9120,
		"delhi":    9145,
		"chennai":  9180,
		"bangalore": 9130,
		"kolkata":  9110,
	}

	price, exists := mockRates[city]
	if !exists {
		return nil, fmt.Errorf("city not supported")
	}

	if karat == "22K" {
		price -= 800
	}

	return &GoldRate{
		Country:      country,
		City:         city,
		Karat:        karat,
		Currency:     "INR",
		PricePerGram: price,
		Source:       "india-provider",
	}, nil
}