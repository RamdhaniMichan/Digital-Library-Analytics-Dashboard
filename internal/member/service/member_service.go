package service

import (
	"digital-library-dashboard/internal/member/model"
	"digital-library-dashboard/internal/member/repository"
	"time"
)

type Service interface {
	Create(m *model.Member) error
	GetByID(id int) (*model.Member, error)
	List() ([]*model.Member, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(m *model.Member) error {
	m.JoinedDate = time.Now()
	return s.repo.Create(m)
}

func (s *service) GetByID(id int) (*model.Member, error) {
	return s.repo.GetByID(id)
}

func (s *service) List() ([]*model.Member, error) {
	return s.repo.List()
}
