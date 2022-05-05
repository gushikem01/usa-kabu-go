package apiconf

import "os"

// FinancialmodelingprepConfig 設定
type FinancialmodelingprepConfig struct {
	APIKey string
}

// NewFinancialmodelingprepConfig ...
func NewFinancialmodelingprepConfig() *FinancialmodelingprepConfig {
	return &FinancialmodelingprepConfig{
		APIKey: getApiKey(),
	}
}

// getApiKey apikey取得
func getApiKey() string {
	return os.Getenv("APIKEY_FINANCIALMODELINGPREP")
}
