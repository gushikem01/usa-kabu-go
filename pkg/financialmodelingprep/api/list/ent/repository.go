package ent

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
	financialmodelingprep "github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/api/list"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
)

type repo struct {
	client *ent.Client
}

const (
	insertMaxCount = 1000
)

// NewRepository リポジトリ作成
func NewRepository(client *ent.Client) financialmodelingprep.Repository {
	return &repo{client}
}

// FindAll データ取得
func (repo *repo) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks := repo.client.Stocks.Query().AllX(ctx)
	return stocks, nil
}

// Create 作成
func (repo *repo) Create(ctx context.Context, stocks []*stocks.Stocks) error {
	bulk := make([]*ent.StocksCreate, insertMaxCount)
	j := 0
	for _, s := range stocks {
		bulk[j] = repo.client.Stocks.Create().SetPrice(s.Price).SetSymbol(s.Symbol).SetExchange(s.Exchange).SetType(s.Type).SetExchange(s.Exchange).SetExchangeShortName(s.ExchangeShortName)
		j++
		if j >= insertMaxCount {
			// create
			_, err := repo.client.Stocks.CreateBulk(bulk...).Save(ctx)
			if err != nil {
				return err
			}
			j = 0
		}
	}
	if j > 0 {
		slice := make([]*ent.StocksCreate, j-1)
		for i := 0; i < j-1; i++ {
			slice[i] = bulk[i]
		}
		_, err := repo.client.Stocks.CreateBulk(slice...).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// Update 更新
func (repo *repo) Update(ctx context.Context, stocks []*stocks.Stocks) error {
	return nil
}
