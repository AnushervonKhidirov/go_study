package repository

import "gorm.io/gorm"

type Repository struct {
	User *UserRepository
	Task *TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Task: NewTaskRepository(db),
	}
}
