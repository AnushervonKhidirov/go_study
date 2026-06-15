package handler

import "task_tracker/internal/service"

type Handler struct {
	User *UserHandler
	Task *TaskHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(services.User),
		Task: NewTaskHandler(services.Task),
	}
}
