package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gushikem01/usa-kabu-go/server/graph/generated"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log *zap.Logger
}

func NewSchema(
	log *zap.Logger,
) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(
		generated.Config{Resolvers: &Resolver{
			log,
		},
		})
}
