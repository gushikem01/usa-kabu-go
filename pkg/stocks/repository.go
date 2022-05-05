package stocks

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
)

// Repository リポジトリ一覧
type Repository interface {
	FindAll(ctx context.Context) ([]*ent.Stocks, error)
}
