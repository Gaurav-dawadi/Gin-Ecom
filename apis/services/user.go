package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us UserService) GetAllUsers() ([]models.User, error) {
	return us.userRepository.GetAllUsers()
}

func (us UserService) CreateUser(user models.User) error {
	return us.userRepository.CreateUser(user)
}

func (us UserService) GetUser(user_id int64) (*models.User, error) {
	return us.userRepository.GetUser(user_id)
}

func (us UserService) GetUserFromEmail(user_email string) (*models.User, error) {
	return us.userRepository.GetUserFromEmail(user_email)
}
