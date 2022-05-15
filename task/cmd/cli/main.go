package main

import (
	"context"
	"os"

	"github.com/gushikem01/usa-kabu-go/pkg/db/dbent"
	"github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/apiconf"
	financialmodelingprep "github.com/gushikem01/usa-kabu-go/pkg/financialmodelingprep/list"
	"github.com/gushikem01/usa-kabu-go/pkg/httpclient"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	stocksEnt "github.com/gushikem01/usa-kabu-go/pkg/stocks/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/zaplog"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	_ "github.com/lib/pq"
)

func initialize() (*financialmodelingprep.Service, *ent.Client, error) {
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

func main() {
	ctx := context.Background()
	sSrv, client, err := initialize()
	defer client.Close()
	if err != nil {
		defer os.Exit(1)
		return
	}
	sSrv.Run(ctx)
}
