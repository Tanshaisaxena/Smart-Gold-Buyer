package providers

import (
	"encoding/json"
	"net/http"
)

type frankfurterResponse struct {
	Amount float64 `json:"amount"`
	Base   string  `json:"base"`
	Date   string  `json:"date"`

	Rates struct {
		INR float64 `json:"INR"`
		EUR float64 `json:"EUR"`
		AED float64 `json:"AED"`
	} `json:"rates"`
}

func FetchUSDINR() (float64, string, error) {

	url := "https://api.frankfurter.app/latest?from=USD&to=INR"

	resp, err := http.Get(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	var result frankfurterResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, "", err
	}

	return result.Rates.INR, result.Date, nil
}
