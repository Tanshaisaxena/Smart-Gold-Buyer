package fetcher

type MarketSnapshot struct {
	Status   string
	Currency string
	Unit     string

	Metals MetalsData

	Timestamps TimestampData
}

type MetalsData struct {
	Gold          float64
	Silver        float64

	LBMAGoldAM   float64
	LBMAGoldPM   float64

	MCXGold      float64
	MCXGoldAM    float64
	MCXGoldPM    float64

	IBJAGold     float64
}

type TimestampData struct {
	Metal    string
	Currency string
}