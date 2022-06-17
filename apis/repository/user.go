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
	err := infrastructure.SetupDatabase().Create(&user).Error
	return err
}

func GetUser(user_id int64) (*models.User, error) {
	user := models.User{}
	err := infrastructure.SetupDatabase().First(&user, user_id).Error
	return &user, err
}

func GetUserFromEmail(user_email string) (*models.User, error) {
	user := models.User{}
	// err := infrastructure.SetupDatabase().Raw("SELECT * FROM users WHERE email = ?", user_email).Scan(&user).Error
	err := infrastructure.SetupDatabase().Where(&models.User{Email: user_email}).Find(&user).Error
	return &user, err
}
