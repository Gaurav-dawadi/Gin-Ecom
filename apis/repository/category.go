package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

type CategoryRepository struct {
	database infrastructure.DatabaseSetup
}

func NewCategoryRepository(database infrastructure.DatabaseSetup) *CategoryRepository {
	return &CategoryRepository{
		database: database,
	}
}

func (cr CategoryRepository) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	err := cr.database.SetupDatabase().Find(&categories).Error
	return categories, err
}

func (cr CategoryRepository) CreateCategory(category models.Category) error {
	err := cr.database.SetupDatabase().Create(&category).Error
	return err
}
