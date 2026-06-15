package service

import (
	"task_tracker/internal/model"
	"task_tracker/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) GetAll() ([]model.User, error) {
	users, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetById(id int) (*model.User, error) {
	user, err := s.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Create(u *model.User) error {
	err := s.repository.Create(u)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(id int, u *model.User) error {
	err := s.repository.Update(id, u)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
