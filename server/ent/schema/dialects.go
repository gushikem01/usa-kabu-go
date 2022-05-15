package schema

import "entgo.io/ent/dialect"

var smallIntSchema = map[string]string{
	dialect.Postgres: "smallint",
}

var bigIntSchema = map[string]string{
	dialect.Postgres: "bigint",
}

var intSchema = map[string]string{
	dialect.Postgres: "int",
}
var floatSchema = map[string]string{
	dialect.Postgres: "float",
}

var textSchema = map[string]string{
	dialect.Postgres: "text",
}
var stringSchema = map[string]string{
	dialect.Postgres: "varchar(255)",
}

var timestampSchema = map[string]string{
	dialect.Postgres: "timestamp",
}
