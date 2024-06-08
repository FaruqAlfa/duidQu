package service

import (
	"duidQu/repository"
	"duidQu/model"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}