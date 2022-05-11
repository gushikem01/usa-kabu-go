package profile_test

import (
	"context"
	"log"
	"testing"

	"github.com/gushikem01/usa-kabu-go/pkg/db/dbent"
	financialmodelingprep "github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/api/profile"
	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	stocksEnt "github.com/gushikem01/usa-kabu-go/pkg/stocks/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/zaplog"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	"github.com/stretchr/testify/assert"
)

// initTest 初期化
func initTest() (*financialmodelingprep.Service, *ent.Client, error) {
	l, e := zaplog.NewZap()
	if e != nil {
		return nil, nil, e
	}

	client, err := dbent.New()
	if err != nil {
		return nil, nil, err
	}
	httpclient := httpclient.NewHTTP(l)
	apiconf := apiconf.NewFinancialmodelingprepConfig()
	sRepo := stocksEnt.NewRepository(client, l)
	sSrv := stocks.NewService(l, sRepo, client)
	fSrv := financialmodelingprep.NewService(l, client, httpclient, sSrv, apiconf)
	return fSrv, client, err
}

// TestRun 実行テスト
func TestRun(t *testing.T) {
	sSrv, client, err := initTest()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	tests := []struct {
		name string
		ctx  context.Context
	}{
		{
			name: "Runテスト",
			ctx:  context.Background(),
		},
	}
	defer client.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := sSrv.Run(tt.ctx)
			assert.Nil(t, e)
		})
	}
}
