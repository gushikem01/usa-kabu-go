package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gushikem01/usa-kabu-go/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	model "github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	endpoint string = "https://financialmodelingprep.com/api/v3/stock/list"
)

// Service サービス
type Service struct {
	logger     *zap.Logger
	repo       Repository
	client     *ent.Client
	httpclient *httpclient.HTTP
	sSrv       *stocks.Service
	apiconf    *apiconf.FinancialmodelingprepConfig
}

// NewService サービス作成
func NewService(logger *zap.Logger, repo Repository, client *ent.Client, httpclient *httpclient.HTTP, sSrv *stocks.Service, apiconf *apiconf.FinancialmodelingprepConfig) *Service {
	return &Service{logger, repo, client, httpclient, sSrv, apiconf}
}

// Run 実行
func (srv *Service) Run(ctx context.Context) error {
	eg := errgroup.Group{}

	var stocks []*ent.Stocks
	eg.Go(func() (err error) {
		stocks, err = srv.FindAll(ctx)
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

	_, _, err := srv.run(ctx, stocks, apiData)
	if err != nil {
		return err
	}
	return nil
}

// run 実行
func (srv *Service) run(ctx context.Context, stocks []*ent.Stocks, apiData []*SymbolsList) (add []*SymbolsList, up []*SymbolsList, err error) {
	symbolMap := map[string]string{}
	for _, v := range stocks {
		symbolMap[v.Symbol] = v.Symbol
	}
	var i, j int
	// 更新
	up = make([]*SymbolsList, len(apiData))

	// 新規
	add = make([]*SymbolsList, len(apiData))

	for _, a := range apiData {
		if v, ok := symbolMap[a.Symbol]; ok {
			up[i] = a
			i++
			fmt.Println(v)
			// Found. update this data
			continue
		}
		// Not Found. insert this data
		add[j] = a
		j++
	}
	return add, up, nil
}

// FindAll 一覧取得
func (srv *Service) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks, err := srv.sSrv.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

// APIGet api実行
func (srv *Service) APIGet(ctx context.Context) ([]*SymbolsList, error) {
	// url := fmt.Sprintf("%s?apikey=%s", endpoint, srv.apiconf.APIKey)
	// fmt.Println(url)
	res, err := srv.httpclient.Get(fmt.Sprintf("%s?apikey=%s", endpoint, srv.apiconf.APIKey), nil)
	if err != nil {
		return nil, err
	}
	jsonBytes := ([]byte)(res)

	var data []*SymbolsList
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// Create 作成
func (srv *Service) Create(ctx context.Context, stocks []*model.Stocks) error {
	if err := srv.repo.Create(ctx, stocks); err != nil {
		return err
	}
	return nil
}
