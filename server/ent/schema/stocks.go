package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Stocks struct {
	ent.Schema
}

// Fields of the Stocks.
func (Stocks) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").SchemaType(stringSchema).Optional(),
		field.String("name").SchemaType(stringSchema).Optional().Default(""),
		field.String("name_ja").SchemaType(stringSchema).Optional().Default(""),
		field.String("type").SchemaType(stringSchema).Optional().Default(""),
		field.String("exchange").SchemaType(stringSchema).Optional().Default(""),
		field.String("exchange_short_name").SchemaType(stringSchema).Optional().Default(""),
		field.Float("price").SchemaType(floatSchema).Optional().Default(0),
		field.String("industry").SchemaType(stringSchema).Optional().Default(""),
		field.Float("market_carp").SchemaType(bigIntSchema).Optional().Default(0),
		field.Float("last_div").SchemaType(floatSchema).Optional().Default(0),
		field.Float("description").SchemaType(textSchema).Optional().Default(0),
		field.String("website").SchemaType(stringSchema).Optional().Default(""),
		field.String("ceo").SchemaType(stringSchema).Optional().Default(""),
		field.String("sector").SchemaType(stringSchema).Optional().Default(""),
		field.String("country").SchemaType(stringSchema).Optional().Default(""),
		field.String("full_time_employees").SchemaType(stringSchema).Optional().Default(""),
		field.String("phone").SchemaType(stringSchema).Optional().Default(""),
		field.String("address").SchemaType(stringSchema).Optional().Default(""),
		field.String("city").SchemaType(stringSchema).Optional().Default(""),
		field.String("state").SchemaType(stringSchema).Optional().Default(""),
		field.String("zip").SchemaType(stringSchema).Optional().Default(""),
		field.String("image").SchemaType(stringSchema).Optional().Default(""),
		field.String("ipo_date").SchemaType(stringSchema).Optional().Default(""),
		field.Float("yield").SchemaType(floatSchema).Optional().Default(0),
		field.Time("created_at").SchemaType(timestampSchema).Default(time.Now),
		field.Time("updated_at").SchemaType(timestampSchema).Default(time.Now),
	}
}

// Edges of the Stocks.
func (Stocks) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stocks", Stocks.Type),
	}
}
