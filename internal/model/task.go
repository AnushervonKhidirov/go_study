package model

type Task struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}
