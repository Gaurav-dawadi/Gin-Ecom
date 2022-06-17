package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(user models.User) error {
	return repository.CreateUser(user)
}

func GetUser(user_id int64) (*models.User, error) {
	return repository.GetUser(user_id)
}

func GetUserFromEmail(user_email string) (*models.User, error) {
	return repository.GetUserFromEmail(user_email)
}
