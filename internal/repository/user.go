package repository

import (
	"context"
	"task_tracker/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := gorm.G[model.User](r.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetById(id int, ctx context.Context) (*model.User, error) {
	user, err := gorm.G[model.User](r.db).Where(getByIdQuery, id).First(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(u *model.User, ctx context.Context) error {
	err := gorm.G[model.User](r.db).Create(ctx, u)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(id int, u *model.User, ctx context.Context) error {
	_, err := gorm.G[model.User](r.db).Where(getByIdQuery, id).Updates(ctx, *u)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int, ctx context.Context) error {
	_, err := gorm.G[model.User](r.db).Where(getByIdQuery, id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
