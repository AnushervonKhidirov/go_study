package model

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUser struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}
