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

func (r *UserRepository) GetAll() (*[]model.User, error) {
	ctx := context.Background()
	users, err := gorm.G[model.User](r.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	ctx := context.Background()
	user, err := gorm.G[model.User](r.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(u *model.User) error {
	ctx := context.Background()
	err := gorm.G[model.User](r.db).Create(ctx, u)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(id int, u *model.User) error {
	ctx := context.Background()
	_, err := gorm.G[model.User](r.db).Where("id = ?", id).Updates(ctx, *u)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	ctx := context.Background()
	_, err := gorm.G[model.User](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
