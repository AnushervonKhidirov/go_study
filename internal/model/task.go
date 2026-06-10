package model

type Task struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

type CreateTask struct {
	Title     string `validate:"required"`
	Desc      string `validate:"required"`
	Completed bool   `validate:"required,boolean"`
}
