package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := infrastructure.SetupDatabase().Find(&users).Error
	return users, err
}

func CreateUser(user models.User) error {
	err := infrastructure.SetupDatabase().Create(user).Error
	return err
}
