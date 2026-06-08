package service

import "fmt"

func GetAllUsers() string {
	return "all_users"
}

func GetSingleUser(id uint) string {
	return fmt.Sprintf("single_user_%d", id)
}
