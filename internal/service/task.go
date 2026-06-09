package service

import (
	"task_tracker/internal/model"
)

type TaskService struct{}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	tasks := []model.Task{
		model.Task{Id: 1, Title: "task title 1", Desc: "task desc 1", Completed: false},
		model.Task{Id: 2, Title: "task title 2", Desc: "task desc 2", Completed: false},
		model.Task{Id: 3, Title: "task title 3", Desc: "task desc 3", Completed: true},
		model.Task{Id: 4, Title: "task title 4", Desc: "task desc 4", Completed: false},
	}
	return tasks, nil
}

func (s *TaskService) GetSingleTask(id uint) (*model.Task, error) {
	task := model.Task{Id: id, Title: "task title 1", Desc: "task desc 1", Completed: false}
	return &task, nil
}
