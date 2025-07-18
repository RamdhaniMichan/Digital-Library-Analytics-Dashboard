package service

import (
	"digital-library-dashboard/internal/analytics/model"
	"digital-library-dashboard/internal/analytics/repository"
)

type AnalyticsService interface {
	GetAnalytics() (model.AnalyticsResponse, error)
}

type analyticsService struct {
	repo repository.AnalyticsRepository
}

func NewAnalyticsService(repo repository.AnalyticsRepository) AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *analyticsService) GetAnalytics() (model.AnalyticsResponse, error) {
	return s.repo.GetAnalytics()
}
