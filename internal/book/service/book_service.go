package service

import (
	"digital-library-dashboard/internal/book/model"
	"digital-library-dashboard/internal/book/repository"
	"digital-library-dashboard/pkg/utils"
)

type Service interface {
	GetAll(page, limit int) ([]model.BookWithCategory, *utils.Paginate, error)
	GetByID(id int) (*model.BookWithCategory, error)
	Create(b model.Book) error
	Update(b model.Book) error
	Delete(id int) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(page, limit int) ([]model.BookWithCategory, *utils.Paginate, error) {
	books, totalItems, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, nil, err
	}
	paginate := utils.NewPaginate(page, limit, totalItems)
	return books, paginate, nil
}

func (s *service) GetByID(id int) (*model.BookWithCategory, error) {
	return s.repo.GetByID(id)
}

func (s *service) Create(b model.Book) error {
	return s.repo.Create(b)
}

func (s *service) Update(b model.Book) error {
	return s.repo.Update(b)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
