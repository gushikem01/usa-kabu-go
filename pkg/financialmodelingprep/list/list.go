package list

import (
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
)

// SymbolsList This is a api model of https://site.financialmodelingprep.com/developer/docs#Symbols-List
type SymbolsList struct {
	ID                int64   `bun:",pk,autoincrement"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Symbol            string  `json:"symbol"`
	Type              string  `json:"type"`
}

// CopyToStocks 構造体コピー
func CopyToStocks(s []*SymbolsList) []*stocks.Stocks {
	st := make([]*stocks.Stocks, len(s))
	for i, v := range s {
		stock := &stocks.Stocks{
			ID:                v.ID,
			Exchange:          v.Exchange,
			ExchangeShortName: v.ExchangeShortName,
			Name:              v.Name,
			Price:             v.Price,
			Symbol:            v.Symbol,
			Type:              v.Type,
		}
		st[i] = stock
	}
	return st
}
