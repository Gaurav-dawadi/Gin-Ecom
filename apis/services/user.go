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
