package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/server/graph/generated"
	"github.com/gushikem01/usa-kabu-go/server/graph/model"
)

func (r *queryResolver) Stocks(ctx context.Context, id *string) (*model.Stocks, error) {
	// panic(fmt.Errorf("not implemented"))
	m := &model.Stocks{
		ID:   "1",
		Name: "tests",
	}
	return m, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
