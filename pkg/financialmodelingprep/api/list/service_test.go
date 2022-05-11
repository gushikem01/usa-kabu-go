package list_test

import (
	"context"
	"log"
	"testing"

	"github.com/gushikem01/usa-kabu-go/pkg/db/dbent"
	financialmodelingprep "github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/api/list"
	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	stocksEnt "github.com/gushikem01/usa-kabu-go/pkg/stocks/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/zaplog"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	_ "github.com/lib/pq"
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

// TestFindAll 一覧テスト
func TestFindAll(t *testing.T) {
	sSrv, client, err := initTest()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	tests := []struct {
		name string
		ctx  context.Context
	}{
		{
			name: "FindAllテスト",
			ctx:  context.Background(),
		},
	}
	defer client.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// e := sSrv.Run(tt.ctx)
			s, e := sSrv.FindAll(tt.ctx)
			assert.NotNil(t, s)
			assert.Nil(t, e)
		})
	}
}

// TestApiTest API実行テスト
func TestApiGet(t *testing.T) {
	sSrv, client, err := initTest()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	tests := []struct {
		name string
		ctx  context.Context
	}{
		{
			name: "Getテスト",
			ctx:  context.Background(),
		},
	}
	defer client.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, e := sSrv.APIGet(tt.ctx)
			assert.NotNil(t, s)
			assert.Nil(t, e)
		})
	}
}
