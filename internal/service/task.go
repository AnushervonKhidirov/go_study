package service

import (
	"task_tracker/internal/model"
	"task_tracker/internal/repository"
)

type TaskService struct {
	repository *repository.TaskRepository
}

func NewTaskService(r *repository.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	tasks, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (s *TaskService) GetById(id int) (*model.Task, error) {
	task, err := s.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return task, err
}

func (s *TaskService) Create(t *model.Task) error {
	err := s.repository.Create(t)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Update(id int, t *model.Task) error {
	err := s.repository.Update(id, t)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
