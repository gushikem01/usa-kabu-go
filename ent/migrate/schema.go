// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StocksColumns holds the columns for the "stocks" table.
	StocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symbol", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "name_ja", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "type", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(16)"}},
		{Name: "exchange", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "exchange_short_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(64)"}},
		{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
	}
	// StocksTable holds the schema information for the "stocks" table.
	StocksTable = &schema.Table{
		Name:       "stocks",
		Columns:    StocksColumns,
		PrimaryKey: []*schema.Column{StocksColumns[0]},
	}
	// StocksStocksColumns holds the columns for the "stocks_stocks" table.
	StocksStocksColumns = []*schema.Column{
		{Name: "stocks_id", Type: field.TypeInt},
		{Name: "stock_id", Type: field.TypeInt},
	}
	// StocksStocksTable holds the schema information for the "stocks_stocks" table.
	StocksStocksTable = &schema.Table{
		Name:       "stocks_stocks",
		Columns:    StocksStocksColumns,
		PrimaryKey: []*schema.Column{StocksStocksColumns[0], StocksStocksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stocks_stocks_stocks_id",
				Columns:    []*schema.Column{StocksStocksColumns[0]},
				RefColumns: []*schema.Column{StocksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "stocks_stocks_stock_id",
				Columns:    []*schema.Column{StocksStocksColumns[1]},
				RefColumns: []*schema.Column{StocksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StocksTable,
		StocksStocksTable,
	}
)

func init() {
	StocksStocksTable.ForeignKeys[0].RefTable = StocksTable
	StocksStocksTable.ForeignKeys[1].RefTable = StocksTable
}
