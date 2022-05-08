package ent

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"go.uber.org/zap"
)

// Repo db構造体
type Repo struct {
	client *ent.Client
	log    *zap.Logger
}

const (
	insertMaxCount = 1000
)

// NewRepository リポジトリ作成
func NewRepository(client *ent.Client, log *zap.Logger) stocks.Repository {
	return &Repo{client, log}
}

// FindAll データ取得
func (repo *Repo) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks := repo.client.Stocks.Query().AllX(ctx)
	return stocks, nil
}

// Create 作成
func (repo *Repo) Create(ctx context.Context, stocks []*stocks.Stocks) error {
	bulk := make([]*ent.StocksCreate, insertMaxCount)
	j := 0
	for _, s := range stocks {
		bulk[j] = repo.client.Stocks.Create().
			SetPrice(s.Price).
			SetSymbol(s.Symbol).
			SetExchange(s.Exchange).
			SetType(s.Type).
			SetExchange(s.Exchange).
			SetExchangeShortName(s.ExchangeShortName)
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

// UpdateAll 更新
func (repo *Repo) UpdateAll(ctx context.Context, stocks []*stocks.Stocks) ([]*ent.Stocks, error) {
	st := make([]*ent.Stocks, len(stocks))
	repo.log.Info(
		"データ更新開始します。",
		zap.Int("件数", len(stocks)),
	)
	for i, s := range stocks {
		n, err := repo.UpdateOne(ctx, s)
		if err != nil {
			return nil, err
		}
		st[i] = n
	}
	repo.log.Info(
		"データ更新完了しました。",
		zap.Int("件数", len(stocks)),
	)

	return st, nil
}

// UpdateOne 1件更新
func (repo *Repo) UpdateOne(ctx context.Context, s *stocks.Stocks) (*ent.Stocks, error) {

	n, err := repo.client.Stocks.UpdateOneID(int(s.ID)).
		SetNillableExchange(&s.Exchange).
		SetNillableExchangeShortName(&s.ExchangeShortName).
		SetNillablePrice(&s.Price).
		SetNillableName(&s.Name).
		SetNillableSymbol(&s.Symbol).
		SetNillableType(&s.Type).
		SetNillableIndustry(&s.Industry).
		SetNillableMarketCarp(&s.MarketCarp).
		Save(ctx)
	if err != nil {
		repo.log.Error(
			"データ更新に失敗しました。",
			zap.Int64("ID", s.ID),
			zap.String("name", s.Name),
			zap.String("symbol", s.Symbol),
			zap.Error(err),
		)
		return nil, err
	}
	// repo.log.Info(
	// 	"データ更新しました。",
	// 	zap.Int64("ID", s.ID),
	// 	zap.String("name", s.Name),
	// 	zap.String("symbol", s.Symbol),
	// )
	return n, nil
}
