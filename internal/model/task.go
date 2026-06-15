package model

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Completed   bool           `json:"completed"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type CreateTask struct {
	Title       string         `validate:"required"`
	Description sql.NullString `validate:"required"`
	Completed   bool           `validate:"required,boolean"`
}
