package service

import (
	"context"
	"digital-library-dashboard/internal/analytics/model"
	"digital-library-dashboard/internal/analytics/repository"
)

type AnalyticsService interface {
	GetAnalytics(ctx context.Context) (model.AnalyticsResponse, error)
}

type analyticsService struct {
	repo repository.AnalyticsRepository
}

func NewAnalyticsService(repo repository.AnalyticsRepository) AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *analyticsService) GetAnalytics(ctx context.Context) (model.AnalyticsResponse, error) {
	return s.repo.GetAnalytics(ctx)
}
