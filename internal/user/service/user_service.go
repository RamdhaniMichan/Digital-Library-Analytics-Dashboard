package service

import (
	"digital-library-dashboard/internal/user/model"
	"digital-library-dashboard/internal/user/repository"
	"digital-library-dashboard/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *model.User) error
	Login(email, password string) (string, string, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(user *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.repo.Create(user)
}

func (s *service) Login(email, password string) (string, string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", err
	}
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	return token, user.Role, nil
}
