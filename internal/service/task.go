package service

import "fmt"

func GetAllTasks() string {
	return "all_tasks"
}

func GetSingleTask(id uint) string {
	return fmt.Sprintf("single_task_%d", id)
}
