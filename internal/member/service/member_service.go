package service

import (
	"digital-library-dashboard/internal/member/model"
	"digital-library-dashboard/internal/member/repository"
	"digital-library-dashboard/pkg/utils"
	"time"
)

type Service interface {
	Create(m *model.Member) error
	GetByID(id int) (*model.Member, error)
	List(page, limit int, filter model.MemberFilter) ([]*model.Member, *utils.Paginate, error)
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

func (s *service) List(page, limit int, filter model.MemberFilter) ([]*model.Member, *utils.Paginate, error) {
	members, totalItems, err := s.repo.List(page, limit, filter)
	if err != nil {
		return nil, nil, err
	}

	paginate := utils.NewPaginate(page, limit, totalItems)
	return members, paginate, nil
}
