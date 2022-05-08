package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gushikem01/usa-kabu-go/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	apiEndpoint   string = "https://financialmodelingprep.com/api/v3/profile/"
	apiMaxRequest int    = 15
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

	stocks, err := srv.FindAll(ctx)
	if err != nil {
		return err
	}
	i := 0
	for _, s := range stocks {
		i++
		sb := s.Symbol
		id := s.ID
		// fmt.Println("件数:", i)
		eg.Go(func() (err error) {
			err = srv.run(ctx, id, sb)
			return
		})
		if i > apiMaxRequest {
			time.Sleep(time.Second * 1)
			i = 0
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// run 実行
func (srv *Service) run(ctx context.Context, id int, symbol string) error {
	p, err := srv.APIGet(ctx, symbol)
	if err != nil {
		srv.log.Error(
			"API取得に失敗しました。",
			zap.String("symbol", symbol),
			zap.Error(err),
		)
		return err
	}
	e := srv.Update(ctx, id, p)
	if e != nil {
		srv.log.Error(
			"更新に失敗しました。",
			zap.String("symbol", symbol),
			zap.Int("ID", id),
			zap.Error(e),
		)
		return e
	}
	srv.log.Info(
		"更新成功しました。",
		zap.String("symbol", symbol),
		zap.Int("ID", id),
	)
	return nil
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
func (srv *Service) APIGet(ctx context.Context, symbol string) (*Profile, error) {
	req := fmt.Sprintf("%s%s?apikey=%s", apiEndpoint, symbol, srv.apiconf.APIKey)
	res, err := srv.httpclient.Get(req, nil)
	if err != nil {
		srv.log.Error(
			"APIの取得に失敗しました。",
			zap.Error(err),
		)
		return nil, err
	}
	jsonBytes := ([]byte)(res)

	var data []*Profile
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	return data[0], nil
}

// Update 更新
func (srv *Service) Update(ctx context.Context, id int, p *Profile) error {
	stock := CopyToStocks(id, p)
	_, err := srv.sSrv.Update(ctx, stock)
	if err != nil {
		return err
	}
	return nil
}
