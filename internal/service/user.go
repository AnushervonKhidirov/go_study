package service

import (
	"task_tracker/internal/model"
)

type UserService struct{}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users := []model.User{
		{Id: 1, Email: "someemail@gmail.com", Name: "Some"},
		{Id: 2, Email: "someemail@yandex.com", Name: "Name"},
	}

	return users, nil
}

func (s *UserService) GetSingleUser(id uint) (*model.User, error) {
	user := model.User{Id: id, Email: "someemail@gmail.com", Name: "Some"}
	return &user, nil
}
