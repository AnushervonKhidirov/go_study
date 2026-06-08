package service

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register()

func GetAllUsers() string {
	return "all_users"
}

func GetSingleUsers() string {
	return "single_user"
}
