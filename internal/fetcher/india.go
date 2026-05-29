package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type apiResponse struct {
	Status   string `json:"status"`
	Currencies CurrencyData `json:"currencies"`
	Unit     string `json:"unit"`

	Metals struct {
		Gold float64 `json:"gold"`

		Silver    float64 `json:"silver"`
		Platinum  float64 `json:"platinum"`
		Palladium float64 `json:"palladium"`

		LBMAGoldAM float64 `json:"lbma_gold_am"`
		LBMAGoldPM float64 `json:"lbma_gold_pm"`

		MCXGold   float64 `json:"mcx_gold"`
		MCXGoldAM float64 `json:"mcx_gold_am"`
		MCXGoldPM float64 `json:"mcx_gold_pm"`

		IBJAGold float64 `json:"ibja_gold"`
	} `json:"metals"`

	Timestamps struct {
		Metal    string `json:"metal"`
		Currency string `json:"currency"`
	} `json:"timestamps"`
}

func FetchMarketSnapshot() (*MarketSnapshot, error) {

	apiKey := os.Getenv("METALS_API_KEY")

	url := fmt.Sprintf(
		"https://api.metals.dev/v1/latest?api_key=%s&currency=INR&unit=g",
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
	return nil, fmt.Errorf("Api returned status %d", resp.StatusCode)
}
	defer resp.Body.Close()

	var apiResp apiResponse

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	snapshot := &MarketSnapshot{
		Status:   apiResp.Status,
		Currencies: apiResp.Currencies,
		Unit:     apiResp.Unit,

		Metals: MetalsData{
			Gold:   apiResp.Metals.Gold,
			Silver: apiResp.Metals.Silver,
			Platinum: apiResp.Metals.Platinum,
			Palladium: apiResp.Metals.Palladium,

			LBMAGoldAM: apiResp.Metals.LBMAGoldAM,
			LBMAGoldPM: apiResp.Metals.LBMAGoldPM,

			MCXGold:   apiResp.Metals.MCXGold,
			MCXGoldAM: apiResp.Metals.MCXGoldAM,
			MCXGoldPM: apiResp.Metals.MCXGoldPM,

			IBJAGold: apiResp.Metals.IBJAGold,
		},

		Timestamps: TimestampData{
			Metal:    apiResp.Timestamps.Metal,
			Currency: apiResp.Timestamps.Currency,
		},
	}

	return snapshot, nil
}
