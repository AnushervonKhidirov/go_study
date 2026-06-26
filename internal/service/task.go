package service

import (
	"context"
	"task_tracker/internal/model"
	"task_tracker/internal/repository"
)

type TaskService struct {
	repository *repository.TaskRepository
}

func NewTaskService(r *repository.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (s *TaskService) GetById(id int, ctx context.Context) (*model.Task, error) {
	task, err := s.repository.GetById(id, ctx)

	if err != nil {
		return nil, err
	}

	return task, err
}

func (s *TaskService) Create(t *model.Task, ctx context.Context) error {
	err := s.repository.Create(t, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Update(id int, t *model.Task, ctx context.Context) error {
	err := s.repository.Update(id, t, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Delete(id int, ctx context.Context) error {
	err := s.repository.Delete(id, ctx)

	if err != nil {
		return err
	}

	return nil
}
