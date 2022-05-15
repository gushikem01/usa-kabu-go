package profile

import "github.com/gushikem01/usa-kabu-go/pkg/stocks"

type Profile struct {
	ID                int64
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int     `json:"volAvg"`
	MktCap            float64 `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	Cik               string  `json:"cik"`
	Isin              string  `json:"isin"`
	Cusip             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	Ceo               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DcfDiff           float64 `json:"dcfDiff"`
	Dcf               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IpoDate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// CopyToStocks 構造体コピー
func CopyToStocks(id int, p *Profile) *stocks.Stocks {
	stock := &stocks.Stocks{
		ID:         int64(id),
		Symbol:     p.Symbol,
		Industry:   p.Industry,
		MarketCarp: p.MktCap,
	}
	return stock
}
