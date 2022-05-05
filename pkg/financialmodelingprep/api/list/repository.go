package list

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
	model "github.com/gushikem01/usa-kabu-go/pkg/stocks"
)

// Repository リポジトリ一覧
type Repository interface {
	FindAll(context.Context) ([]*ent.Stocks, error)
	Create(context.Context, []*model.Stocks) error
	Update(context.Context, []*model.Stocks) error
}
