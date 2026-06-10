package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"task_tracker/internal/model"
	"task_tracker/pkg/apperr"
)

type UserService struct{}

func (s *UserService) GetAllUsers() (*[]model.User, error) {
	userList, err := s.readFromFile()

	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (s *UserService) GetSingleUser(id uint) (*model.User, error) {
	userList, err := s.readFromFile()

	if err != nil {
		return nil, err
	}

	var user *model.User

	for _, u := range *userList {
		if u.Id == id {
			user = &u
			break
		}
	}

	if user == nil {
		message := "User not found"
		return nil, apperr.NotFoundErr(&message)
	}

	return user, nil
}

func (s *UserService) AddUser(u *model.CreateUser) error {
	userList, err := s.readFromFile()

	if err != nil {
		return err
	}

	var lastId uint

	if len(*userList) == 0 {
		lastId = 1
	} else {
		lastId = (*userList)[len(*userList)-1].Id + 1
	}

	*userList = append(*userList, model.User{Id: lastId, Email: u.Email, Name: u.Name})

	err = s.writeToFile(userList)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) readFromFile() (*[]model.User, error) {
	cwd, err := os.Getwd()
	databasePath := filepath.Join(cwd, "database", "users.json")

	if err != nil {
		return nil, err
	}

	database, err := os.ReadFile(databasePath)

	if err != nil {
		return nil, err
	}

	var userList []model.User

	err = json.Unmarshal(database, &userList)

	if err != nil {
		return nil, err
	}

	return &userList, nil
}

func (s *UserService) writeToFile(d *[]model.User) error {
	cwd, err := os.Getwd()
	databasePath := filepath.Join(cwd, "database", "users.json")

	if err != nil {
		return err
	}

	userListJson, err := json.Marshal(&d)

	if err != nil {
		return err
	}

	err = os.WriteFile(databasePath, userListJson, 0666)

	if err != nil {
		return err
	}

	return nil
}
