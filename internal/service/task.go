package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"task_tracker/internal/model"
	"task_tracker/pkg/apperr"
)

type TaskService struct{}

func (s *TaskService) GetAllTasks() (*[]model.Task, error) {
	tasks, err := s.readFromFile()

	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func (s *TaskService) GetSingleTask(id uint) (*model.Task, error) {
	tasks, err := s.readFromFile()

	if err != nil {
		return nil, err
	}

	var task *model.Task

	for _, v := range *tasks {
		if v.Id == id {
			task = &v
		}
	}

	if task == nil {
		var appErr apperr.AppErr
		message := "Task not found"
		return nil, appErr.NotFoundErr(&message)
	}

	return task, nil

}

func (s *TaskService) AddTask(t *model.CreateTask) error {
	taskList, err := s.readFromFile()

	if err != nil {
		return err
	}

	var lastId uint

	if len(*taskList) == 0 {
		lastId = 1
	} else {
		lastId = (*taskList)[len(*taskList)-1].Id + 1
	}

	*taskList = append(*taskList, model.Task{Id: lastId, Title: t.Title, Desc: t.Desc, Completed: t.Completed})

	err = s.writeToFile(taskList)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) readFromFile() (*[]model.Task, error) {
	cwd, err := os.Getwd()
	databasePath := filepath.Join(cwd, "database", "tasks.json")

	if err != nil {
		return nil, err
	}

	database, err := os.ReadFile(databasePath)

	if err != nil {
		return nil, err
	}

	var taskList []model.Task

	err = json.Unmarshal(database, &taskList)

	if err != nil {
		return nil, err
	}

	return &taskList, nil
}

func (s *TaskService) writeToFile(d *[]model.Task) error {
	cwd, err := os.Getwd()
	databasePath := filepath.Join(cwd, "database", "tasks.json")

	if err != nil {
		return err
	}

	taskListJson, err := json.Marshal(&d)

	if err != nil {
		return err
	}

	err = os.WriteFile(databasePath, taskListJson, 0666)

	if err != nil {
		return err
	}

	return nil
}
