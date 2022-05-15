package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	apiEndpoint string = "https://financialmodelingprep.com/api/v3/stock/list"
)

// Service サービス
type Service struct {
	log        *zap.Logger
	client     *ent.Client
	httpclient *httpclient.HTTP
	sSrv       *stocks.Service
	apiconf    *apiconf.FinancialmodelingprepConfig
}

// NewService サービス作成
func NewService(log *zap.Logger, client *ent.Client, httpclient *httpclient.HTTP, sSrv *stocks.Service, apiconf *apiconf.FinancialmodelingprepConfig) *Service {
	return &Service{log, client, httpclient, sSrv, apiconf}
}

// Run 実行
func (srv *Service) Run(ctx context.Context) error {
	eg := errgroup.Group{}

	var stocks []*ent.Stocks
	eg.Go(func() (err error) {
		stocks, err = srv.FindAll(ctx)
		srv.log.Info(
			"stocksデータ取得しました。",
			zap.Int("stocks", len(stocks)),
		)
		return
	})

	var apiData []*SymbolsList
	eg.Go(func() (err error) {
		apiData, err = srv.APIGet(ctx)
		return
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	_, err := srv.run(ctx, stocks, apiData)
	if err != nil {
		return err
	}
	return nil
}

// run 実行
func (srv *Service) run(ctx context.Context, stocks []*ent.Stocks, apiData []*SymbolsList) ([]*ent.Stocks, error) {
	var u []*ent.Stocks
	var err error
	add, up := srv.createOrUpdate(ctx, stocks, apiData)

	eg := errgroup.Group{}
	eg.Go(func() (err error) {
		err = srv.Create(ctx, add)
		return
	})
	// TODO: 更新処理はここでは行わない
	_ = up
	// eg.Go(func() (err error) {
	// 	u, err = srv.UpdateAll(ctx, up)
	// 	return
	// })

	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return u, err
}

// createOrUpdate 作成と更新
// 	引数1:	Stocksテーブルのデータ
//	引数2: 	Apiデータ
//	戻り値1:	追加データ
//	戻り値2:	追加データ
func (srv *Service) createOrUpdate(ctx context.Context, stocks []*ent.Stocks, apiData []*SymbolsList) (add []*SymbolsList, up []*SymbolsList) {
	var i, j int
	up = make([]*SymbolsList, len(apiData))

	add = make([]*SymbolsList, len(apiData))

	for _, a := range apiData {
		find := false
		id := 0
		for _, v := range stocks {
			if v.Symbol == a.Symbol {
				id = v.ID
				find = true
				break
			}
		}
		if find {
			a.ID = int64(id)
			up[i] = a
			i++
			continue
		}

		// Not Found. insert this data
		add[j] = a
		j++
	}
	if len(up) > 0 && i > 0 {
		up = up[:i-1]
	}
	if len(add) > 0 && j > 0 {
		add = add[:j-1]
	}
	return add, up
}

// FindAll 一覧取得
func (srv *Service) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks, err := srv.sSrv.FindAll(ctx)
	if err != nil {
		srv.log.Error(
			"一覧取得に失敗しました。",
			zap.Error(err),
		)
		return nil, err
	}
	return stocks, nil
}

// APIGet api実行
func (srv *Service) APIGet(ctx context.Context) ([]*SymbolsList, error) {
	res, err := srv.httpclient.Get(fmt.Sprintf("%s?apikey=%s", apiEndpoint, srv.apiconf.APIKey), nil)
	if err != nil {
		srv.log.Error(
			"APIの取得に失敗しました。",
			zap.String("API", apiEndpoint),
			zap.Error(err),
		)
		return nil, err
	}
	jsonBytes := ([]byte)(res)

	var data []*SymbolsList
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	srv.log.Info(
		"APIの取得に成功しました。",
		zap.String("API", apiEndpoint),
		zap.Int("件数", len(data)),
	)
	return data, nil
}

// Create 作成
func (srv *Service) Create(ctx context.Context, stocks []*SymbolsList) error {
	if len(stocks) == 0 {
		return nil
	}
	s := CopyToStocks(stocks)
	if err := srv.sSrv.Create(ctx, s); err != nil {
		srv.log.Error(
			"データ作成に失敗しました。",
			zap.Error(err),
		)
		return err
	}
	srv.log.Info(
		"データ作成に成功しました。",
		zap.Int("件数", len(stocks)),
	)
	return nil
}

// Update 更新
func (srv *Service) UpdateAll(ctx context.Context, stocks []*SymbolsList) ([]*ent.Stocks, error) {
	if len(stocks) == 0 {
		return nil, nil
	}

	s := CopyToStocks(stocks)
	up, err := srv.sSrv.UpdateAll(ctx, s)
	if err != nil {
		srv.log.Error(
			"データ更新に失敗しました。",
			zap.Error(err),
		)
		return nil, err
	}
	srv.log.Info(
		"データ更新に成功しました。",
		zap.Int("件数", len(stocks)),
	)
	return up, nil
}
