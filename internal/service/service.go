package service

import "task_tracker/internal/repository"

type Service struct {
	User *UserService
	Task *TaskService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Task: NewTaskService(repos.Task),
	}
}
