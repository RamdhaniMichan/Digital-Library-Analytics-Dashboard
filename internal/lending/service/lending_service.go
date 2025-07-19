package service

import (
	bookRepository "digital-library-dashboard/internal/book/repository"
	"digital-library-dashboard/internal/lending/model"
	"digital-library-dashboard/internal/lending/repository"
	memberRepository "digital-library-dashboard/internal/member/repository"
	"digital-library-dashboard/pkg/utils"
	"errors"
)

type LendingService interface {
	Create(l model.Lending) error
	GetAll(page, limit int) ([]model.Lending, *utils.Paginate, error)
	GetByID(id int) (model.Lending, error)
	Update(l model.Lending) error
	Delete(id int) error
}

type lendingService struct {
	repo       repository.LendingRepository
	bookRepo   bookRepository.Repository
	memberRepo memberRepository.Repository
}

func NewLendingService(repo repository.LendingRepository,
	bookRepo bookRepository.Repository,
	memberRepo memberRepository.Repository) LendingService {
	return &lendingService{repo: repo, bookRepo: bookRepo, memberRepo: memberRepo}
}

func (s *lendingService) Create(l model.Lending) error {
	if l.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	if l.BookID <= 0 || l.MemberID <= 0 {
		return errors.New("book ID and member ID must be greater than zero")
	}

	book, err := s.bookRepo.GetByID(l.BookID)
	if err != nil {
		return errors.New("book not found")
	}

	if book.Quantity <= 0 {
		return errors.New("book is not available")
	}

	member, err := s.memberRepo.GetByID(l.MemberID)
	if err != nil {
		return errors.New("member not found")
	}

	if member.Status != "active" {
		return errors.New("member is not active")
	}

	return s.repo.Create(l)
}

func (s *lendingService) GetAll(page, limit int) ([]model.Lending, *utils.Paginate, error) {
	lendings, totalItems, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, nil, err
	}
	paginate := utils.NewPaginate(page, limit, totalItems)
	return lendings, paginate, nil
}

func (s *lendingService) GetByID(id int) (model.Lending, error) {
	return s.repo.GetByID(id)
}

func (s *lendingService) Update(l model.Lending) error {
	lending, err := s.repo.GetByID(l.ID)
	if err != nil {
		return err
	}

	if l.ReturnDate != nil && lending.DueDate.Format("2006-01-02") < l.ReturnDate.Format("2006-01-02") {
		lending.Status = "late"
	} else {
		lending.Status = "returned"
	}

	lending.ReturnDate = l.ReturnDate
	return s.repo.Update(lending)
}

func (s *lendingService) Delete(id int) error {
	return s.repo.Delete(id)
}
