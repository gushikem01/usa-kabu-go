package list

// SymbolsList This is a api model of https://site.financialmodelingprep.com/developer/docs#Symbols-List
type SymbolsList struct {
	ID                int64   `bun:",pk,autoincrement"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	Price             float64 `json:"price"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
}
