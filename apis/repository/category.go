package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	err := infrastructure.SetupDatabase().Find(&categories).Error
	return categories, err
}

func CreateCategory(category models.Category) error {
	err := infrastructure.SetupDatabase().Create(&category).Error
	return err
}
