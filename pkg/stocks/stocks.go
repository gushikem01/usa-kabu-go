package stocks

// Stocks モデル
type Stocks struct {
	ID                int64   `bun:",pk,autoincrement"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	Price             float64 `json:"price"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	MarketCarp        float64 `json:"market_carp"`
}
