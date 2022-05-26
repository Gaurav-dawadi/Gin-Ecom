package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func GetAllCategory() ([]models.Category, error) {
	return repository.GetAllCategory()
}

func CreateCategory(category models.Category) error {
	err := repository.CreateCategory(category)
	return err
}
