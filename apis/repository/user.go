package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

type UserRepository struct {
	database infrastructure.DatabaseSetup
}

func NewUserRepository(database infrastructure.DatabaseSetup) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (ur UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.database.SetupDatabase().Find(&users).Error
	return users, err
}

func (ur UserRepository) CreateUser(user models.User) error {
	err := ur.database.SetupDatabase().Create(&user).Error
	return err
}

func (ur UserRepository) GetUser(user_id int64) (*models.User, error) {
	user := models.User{}
	err := ur.database.SetupDatabase().First(&user, user_id).Error
	return &user, err
}

func (ur UserRepository) GetUserFromEmail(user_email string) (*models.User, error) {
	user := models.User{}
	// err := infrastructure.SetupDatabase().Raw("SELECT * FROM users WHERE email = ?", user_email).Scan(&user).Error
	err := ur.database.SetupDatabase().Where(&models.User{Email: user_email}).Find(&user).Error
	return &user, err
}
