package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	ent1 "github.com/gushikem01/usa-kabu-go/server/ent"
	"github.com/gushikem01/usa-kabu-go/server/graph/generated"
)

func (r *queryResolver) Stocks(ctx context.Context, id *string) ([]*ent1.Stocks, error) {
	// panic(fmt.Errorf("not implemented"))
	s, err := r.sSrv.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *stocksResolver) Name(ctx context.Context, obj *ent1.Stocks) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stocksResolver) Type(ctx context.Context, obj *ent1.Stocks) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stocksResolver) Price(ctx context.Context, obj *ent1.Stocks) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stocksResolver) Exchange(ctx context.Context, obj *ent1.Stocks) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stocksResolver) ExchangeShortName(ctx context.Context, obj *ent1.Stocks) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Stocks returns generated.StocksResolver implementation.
func (r *Resolver) Stocks() generated.StocksResolver { return &stocksResolver{r} }

type queryResolver struct{ *Resolver }
type stocksResolver struct{ *Resolver }
