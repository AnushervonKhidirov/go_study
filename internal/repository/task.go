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

func (r *TaskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks, err := gorm.G[model.Task](r.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(id int, ctx context.Context) (*model.Task, error) {
	task, err := gorm.G[model.Task](r.db).Where(getByIdQuery, id).First(ctx)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Create(t *model.Task, ctx context.Context) error {
	err := gorm.G[model.Task](r.db).Create(ctx, t)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Update(id int, t *model.Task, ctx context.Context) error {
	_, err := gorm.G[model.Task](r.db).Where(getByIdQuery, id).Updates(ctx, *t)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Delete(id int, ctx context.Context) error {
	_, err := gorm.G[model.Task](r.db).Where(getByIdQuery, id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
