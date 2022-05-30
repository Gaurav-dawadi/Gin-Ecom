package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func CreateProductImage(prodImage models.ProductImage) error {
	err := infrastructure.SetupDatabase().Create(&prodImage).Error
	return err
}
