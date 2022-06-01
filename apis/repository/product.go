package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := infrastructure.SetupDatabase().Find(&products).Error
	return products, err
}

func CreateProduct(prod models.Product) (*models.Product, error) {
	err := infrastructure.SetupDatabase().Create(&prod).Error
	return &prod, err
}
