package models

type MarketSnapshot struct {
	Source     string        `json:"source"`
	Status     string        `json:"status"`
	Currency   string        `json:"currency"`
	Unit       string        `json:"unit"`
	Metals     MetalsData    `json:"metals"`
	Currencies CurrencyData  `json:"currencies"`
	Timestamps TimestampData `json:"timestamps"`
}

type MetalsData struct {
	Gold      float64 `json:"gold"`
	Silver    float64 `json:"silver"`
	Platinum  float64 `json:"platinum"`
	Palladium float64 `json:"palladium"`
}

type CurrencyData struct {
	USDINR float64 `json:"usd_inr"`
	EURINR float64 `json:"eur_inr"`
	AEDINR float64 `json:"aed_inr"`
}

type TimestampData struct {
	Metal    string `json:"metal"`
	Currency string `json:"currency"`
}