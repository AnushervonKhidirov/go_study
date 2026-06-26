package service

import (
	"context"
	"task_tracker/internal/model"
	"task_tracker/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetById(id int, ctx context.Context) (*model.User, error) {
	user, err := s.repository.GetById(id, ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Create(u *model.User, ctx context.Context) error {
	err := s.repository.Create(u, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(id int, u *model.User, ctx context.Context) error {
	err := s.repository.Update(id, u, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(id int, ctx context.Context) error {
	err := s.repository.Delete(id, ctx)

	if err != nil {
		return err
	}

	return nil
}
