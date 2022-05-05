package ent

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
)

// Repo db構造体
type Repo struct {
	client *ent.Client
}

// NewRepository リポジトリ作成
func NewRepository(client *ent.Client) stocks.Repository {
	return &Repo{client}
}

// FindAll データ取得
func (repo *Repo) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks := repo.client.Stocks.Query().AllX(ctx)
	return stocks, nil
}
