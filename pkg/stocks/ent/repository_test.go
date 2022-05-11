package ent_test

import (
	"context"
	"log"
	"testing"

	"github.com/gushikem01/usa-kabu-go/pkg/db/dbent"
	"github.com/gushikem01/usa-kabu-go/pkg/stocks"
	stocksEnt "github.com/gushikem01/usa-kabu-go/pkg/stocks/ent"
	"github.com/gushikem01/usa-kabu-go/pkg/zaplog"
	"github.com/gushikem01/usa-kabu-go/server/ent"
	"github.com/stretchr/testify/assert"
)

func initTest() (stocks.Repository, *ent.Client, error) {

	l, e := zaplog.NewZap()
	if e != nil {
		return nil, nil, e
	}

	client, err := dbent.New()
	if err != nil {
		return nil, nil, err
	}
	repo := stocksEnt.NewRepository(client, l)
	return repo, client, nil
}

func TestFindAll(t *testing.T) {
	repo, client, err := initTest()
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
			s, e := repo.FindAll(tt.ctx)
			assert.NotNil(t, s)
			assert.Nil(t, e)
		})
	}

}
