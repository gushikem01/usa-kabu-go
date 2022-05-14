// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/gushikem01/usa-kabu-go/server/ent/schema"
	"github.com/gushikem01/usa-kabu-go/server/ent/stocks"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	stocksFields := schema.Stocks{}.Fields()
	_ = stocksFields
	// stocksDescName is the schema descriptor for name field.
	stocksDescName := stocksFields[1].Descriptor()
	// stocks.DefaultName holds the default value on creation for the name field.
	stocks.DefaultName = stocksDescName.Default.(string)
	// stocksDescNameJa is the schema descriptor for name_ja field.
	stocksDescNameJa := stocksFields[2].Descriptor()
	// stocks.DefaultNameJa holds the default value on creation for the name_ja field.
	stocks.DefaultNameJa = stocksDescNameJa.Default.(string)
	// stocksDescType is the schema descriptor for type field.
	stocksDescType := stocksFields[3].Descriptor()
	// stocks.DefaultType holds the default value on creation for the type field.
	stocks.DefaultType = stocksDescType.Default.(string)
	// stocksDescExchange is the schema descriptor for exchange field.
	stocksDescExchange := stocksFields[4].Descriptor()
	// stocks.DefaultExchange holds the default value on creation for the exchange field.
	stocks.DefaultExchange = stocksDescExchange.Default.(string)
	// stocksDescExchangeShortName is the schema descriptor for exchange_short_name field.
	stocksDescExchangeShortName := stocksFields[5].Descriptor()
	// stocks.DefaultExchangeShortName holds the default value on creation for the exchange_short_name field.
	stocks.DefaultExchangeShortName = stocksDescExchangeShortName.Default.(string)
	// stocksDescPrice is the schema descriptor for price field.
	stocksDescPrice := stocksFields[6].Descriptor()
	// stocks.DefaultPrice holds the default value on creation for the price field.
	stocks.DefaultPrice = stocksDescPrice.Default.(float64)
	// stocksDescIndustry is the schema descriptor for industry field.
	stocksDescIndustry := stocksFields[7].Descriptor()
	// stocks.DefaultIndustry holds the default value on creation for the industry field.
	stocks.DefaultIndustry = stocksDescIndustry.Default.(string)
	// stocksDescMarketCarp is the schema descriptor for market_carp field.
	stocksDescMarketCarp := stocksFields[8].Descriptor()
	// stocks.DefaultMarketCarp holds the default value on creation for the market_carp field.
	stocks.DefaultMarketCarp = stocksDescMarketCarp.Default.(float64)
	// stocksDescLastDiv is the schema descriptor for last_div field.
	stocksDescLastDiv := stocksFields[9].Descriptor()
	// stocks.DefaultLastDiv holds the default value on creation for the last_div field.
	stocks.DefaultLastDiv = stocksDescLastDiv.Default.(int)
	// stocksDescDescription is the schema descriptor for description field.
	stocksDescDescription := stocksFields[10].Descriptor()
	// stocks.DefaultDescription holds the default value on creation for the description field.
	stocks.DefaultDescription = stocksDescDescription.Default.(float64)
	// stocksDescWebsite is the schema descriptor for website field.
	stocksDescWebsite := stocksFields[11].Descriptor()
	// stocks.DefaultWebsite holds the default value on creation for the website field.
	stocks.DefaultWebsite = stocksDescWebsite.Default.(string)
	// stocksDescYield is the schema descriptor for yield field.
	stocksDescYield := stocksFields[12].Descriptor()
	// stocks.DefaultYield holds the default value on creation for the yield field.
	stocks.DefaultYield = stocksDescYield.Default.(float64)
	// stocksDescCreatedAt is the schema descriptor for created_at field.
	stocksDescCreatedAt := stocksFields[13].Descriptor()
	// stocks.DefaultCreatedAt holds the default value on creation for the created_at field.
	stocks.DefaultCreatedAt = stocksDescCreatedAt.Default.(func() time.Time)
	// stocksDescUpdatedAt is the schema descriptor for updated_at field.
	stocksDescUpdatedAt := stocksFields[14].Descriptor()
	// stocks.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	stocks.DefaultUpdatedAt = stocksDescUpdatedAt.Default.(func() time.Time)
}
