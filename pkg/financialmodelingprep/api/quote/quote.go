package quote

type Quote struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Price                float64 `json:"price"`
	ChangesPercentage    float64 `json:"changesPercentage"`
	Change               float64 `json:"change"`
	DayLow               float64 `json:"dayLow"`
	DayHigh              float64 `json:"dayHigh"`
	YearHigh             float64 `json:"yearHigh"`
	YearLow              float64 `json:"yearLow"`
	MarketCap            float64 `json:"marketCap"`
	PriceAvg50           float64 `json:"priceAvg50"`
	PriceAvg200          float64 `json:"priceAvg200"`
	Volume               int     `json:"volume"`
	AvgVolume            int     `json:"avgVolume"`
	Exchange             string  `json:"exchange"`
	Open                 float64 `json:"open"`
	PreviousClose        float64 `json:"previousClose"`
	Eps                  float64 `json:"eps"`
	Pe                   float64 `json:"pe"`
	EarningsAnnouncement string  `json:"earningsAnnouncement"`
	SharesOutstanding    int64   `json:"sharesOutstanding"`
	Timestamp            int     `json:"timestamp"`
}
