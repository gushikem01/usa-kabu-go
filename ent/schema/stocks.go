package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stocks holds the schema definition for the Stocks entity.
type Stocks struct {
	ent.Schema
}

// Fields of the Stocks.
func (Stocks) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.String("name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}).Optional(),
		field.String("name_ja").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}).Optional(),
		field.String("type").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(16)",
			}),
		field.String("exchange").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.String("exchange_short_name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(64)",
			}),
		field.Float("price").
			SchemaType(map[string]string{
				dialect.Postgres: "numeric",
			}),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).Default(time.Now),
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).Default(time.Now),
	}
}

// Edges of the Stocks.
func (Stocks) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stocks", Stocks.Type),
	}
}
