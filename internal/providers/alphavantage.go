package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"Gold-Rate-Analyser/internal/models"
)

type AlphaVantageProvider struct{}

func (p AlphaVantageProvider) Name() string {
	return "alphavantage"
}

type globalQuoteResponse struct {
	GlobalQuote struct {
		Price string `json:"05. price"`
	} `json:"Global Quote"`

	Information string `json:"Information"`
	Note        string `json:"Note"`
}

func fetchQuote(symbol string, apiKey string) (float64, error) {

	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s",
		symbol,
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var quote globalQuoteResponse

	err = json.Unmarshal(body, &quote)
	if err != nil {
		return 0, err
	}

	if quote.Information != "" {
		return 0, fmt.Errorf(
			"alphavantage information response for %s: %s",
			symbol,
			quote.Information,
		)
	}

	if quote.Note != "" {
		return 0, fmt.Errorf(
			"alphavantage rate limit for %s: %s",
			symbol,
			quote.Note,
		)
	}

	if quote.GlobalQuote.Price == "" {
		return 0, fmt.Errorf(
			"price missing for symbol %s",
			symbol,
		)
	}

	return strconv.ParseFloat(
		quote.GlobalQuote.Price,
		64,
	)
}

func fetchQuoteWithDelay(
	symbol string,
	apiKey string,
) (float64, error) {

	price, err := fetchQuote(symbol, apiKey)

	time.Sleep(3 * time.Second)

	return price, err
}

func (p AlphaVantageProvider) FetchMarketSnapshot() (*models.MarketSnapshot, error) {

	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("METALS_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("ALPHA_VANTAGE_API_KEY or METALS_API_KEY not set")
	}

	snapshot := &models.MarketSnapshot{
		Source: "alphavantage",

		Status:   "success",
		Currency: "USD",
		Unit:     "ETF",
	}

	gold, err := fetchQuoteWithDelay("GLD", apiKey)
	if err != nil {
		return nil, fmt.Errorf(
			"gold fetch failed: %w",
			err,
		)
	}
	snapshot.Metals.Gold = gold

	silver, err := fetchQuoteWithDelay("SLV", apiKey)
	if err == nil {
		snapshot.Metals.Silver = silver
	}

	platinum, err := fetchQuoteWithDelay("PPLT", apiKey)
	if err == nil {
		snapshot.Metals.Platinum = platinum
	}

	palladium, err := fetchQuoteWithDelay("PALL", apiKey)
	if err == nil {
		snapshot.Metals.Palladium = palladium
	}

	snapshot.Timestamps.Metal = time.Now().UTC().Format(time.RFC3339)

	usdInr, fxDate, err := FetchUSDINR()
	if err == nil {
		snapshot.Currencies.USDINR = usdInr
		snapshot.Timestamps.Currency = fxDate
	}
	return snapshot, nil
}
