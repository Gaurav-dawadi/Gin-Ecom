package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func GetAllProductImage() ([]models.ProductImage, error) {
	var productImgs []models.ProductImage
	err := infrastructure.SetupDatabase().Find(&productImgs).Error
	return productImgs, err
}

func CreateProductImage(prodImage models.ProductImage) error {
	err := infrastructure.SetupDatabase().Create(&prodImage).Error
	return err
}
