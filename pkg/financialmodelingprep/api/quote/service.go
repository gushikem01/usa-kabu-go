package quote

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	apiEndpoint string = "https://financialmodelingprep.com/api/v3/quote/"
	apiMaxLimit int    = 100
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

	j := 0
	apis := make([]string, apiMaxLimit)

	for _, s := range stocks {
		// 100件ずつAPI実行とDBUpdate
		if j >= apiMaxLimit {
			apis = make([]string, apiMaxLimit)
			j = 0
			eg.Go(func() (err error) {
				err = srv.run(ctx, apis)
				return
			})
		}
		apis[j] = s.Symbol
		j++
	}
	// 残り分をAPI実行とDBUpdate
	if j > 0 {
		apis = apis[:j-1]
		eg.Go(func() (err error) {
			err = srv.run(ctx, apis)
			return
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// run 実行
func (srv *Service) run(ctx context.Context, apis []string) error {
	str := strings.Join(apis, ",")
	p, err := srv.APIGet(ctx, &str)
	if err != nil {
		return err
	}
	fmt.Println(p)
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
func (srv *Service) APIGet(ctx context.Context, str *string) ([]*Quote, error) {
	res, err := srv.httpclient.Get(fmt.Sprintf("%s?apikey=%s", apiEndpoint, srv.apiconf.APIKey), nil)
	if err != nil {
		srv.log.Error(
			"APIの取得に失敗しました。",
			zap.Error(err),
		)
		return nil, err
	}
	jsonBytes := ([]byte)(res)

	var data []*Quote
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}
	return data, nil

}
