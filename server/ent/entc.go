//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// ex, err := entgql.NewExtension(
	// 	entgql.WithWhereFilters(true),
	// 	entgql.WithSchemaPath("../graph/schema/ent.graphql"),
	// 	entgql.WithConfigPath("../gqlgen.yml"),
	// )
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	// opts := []entc.Option{
	// 	entc.Extensions(entviz.Extension{}),
	// 	entc.Extensions(ex),
	// 	entc.TemplateDir("./template"),
	// }
	opts := []entc.Option{
		entc.Extensions(ex),
	}
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureLock,
			gen.FeatureUpsert,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
