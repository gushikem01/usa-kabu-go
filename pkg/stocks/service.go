package stocks

import (
	"context"

	"github.com/gushikem01/usa-kabu-go/ent"
	"go.uber.org/zap"
)

// Service サービス
type Service struct {
	logger *zap.Logger
	repo   Repository
	client *ent.Client
}

// NewService サービス作成
func NewService(logger *zap.Logger, repo Repository, client *ent.Client) *Service {
	return &Service{logger, repo, client}
}

// FindAll 一覧
func (srv *Service) FindAll(ctx context.Context) ([]*ent.Stocks, error) {
	stocks, err := srv.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
