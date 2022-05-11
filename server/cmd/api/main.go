package main

import (
	"github.com/gushikem01/usa-kabu-go/pkg/db/dbent"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/zaplog"
	"github.com/gushikem01/usa-kabu-go/server/graph/resolver"
	"github.com/gushikem01/usa-kabu-go/server/internal/appgin"
	"github.com/gushikem01/usa-kabu-go/server/internal/gql/gqlgin"
)

func main() {
	l, err := zaplog.NewZap()
	client, err := dbent.New()
	repo := ent.NewRepository(client, l)
	sSrv := *stocks.NewService(l, repo, client)
	if err != nil {
		panic(err)
	}
	executableSchema := resolver.NewSchema(l, &sSrv)
	router := gqlgin.NewRouter(l)
	api := appgin.NewAPI(executableSchema, router)
	err = api.Router.Run()
	if err != nil {
		panic(err)
	}
}
