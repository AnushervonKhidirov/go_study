package model

type User struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUser struct {
	Email string
	Name  string
}
