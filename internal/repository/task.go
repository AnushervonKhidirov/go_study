package repository

import (
	"context"
	"task_tracker/internal/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	ctx := context.Background()
	tasks, err := gorm.G[model.Task](r.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(id int) (*model.Task, error) {
	ctx := context.Background()
	task, err := gorm.G[model.Task](r.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Create(t *model.Task) error {
	ctx := context.Background()

	err := gorm.G[model.Task](r.db).Create(ctx, t)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Update(id int, t *model.Task) error {
	ctx := context.Background()
	_, err := gorm.G[model.Task](r.db).Where("id = ?", id).Updates(ctx, *t)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Delete(id int) error {
	ctx := context.Background()
	_, err := gorm.G[model.Task](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
