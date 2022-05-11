package schema

import "entgo.io/ent/dialect"

var smallIntSchema = map[string]string{
	dialect.Postgres: "smallint",
}

var float64Schema = map[string]string{
	dialect.Postgres: "float64",
}

var stringSchema = map[string]string{
	dialect.Postgres: "varchar(255)",
}
