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

// UpdateAll 作成
func (srv *Service) UpdateAll(ctx context.Context, stocks []*Stocks) ([]*ent.Stocks, error) {
	up, err := srv.repo.UpdateAll(ctx, stocks)
	if err != nil {
		return nil, err
	}
	return up, nil
}

// Create 作成
func (srv *Service) Create(ctx context.Context, stocks []*Stocks) error {
	if err := srv.repo.Create(ctx, stocks); err != nil {
		return err
	}
	return nil
}

// Update 更新
func (srv *Service) Update(ctx context.Context, stocks *Stocks) (*ent.Stocks, error) {
	s, err := srv.repo.UpdateOne(ctx, stocks)
	if err != nil {
		return nil, err
	}
	return s, nil
}
