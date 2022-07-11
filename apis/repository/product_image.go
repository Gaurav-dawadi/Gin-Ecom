package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

type ProductImageRepository struct {
	database infrastructure.DatabaseSetup
}

func NewProductImageRepository(database infrastructure.DatabaseSetup) *ProductImageRepository {
	return &ProductImageRepository{
		database: database,
	}
}

func (pir ProductImageRepository) GetAllProductImage() ([]models.ProductImage, error) {
	var productImgs []models.ProductImage
	err := pir.database.SetupDatabase().Find(&productImgs).Error
	return productImgs, err
}

func (pir ProductImageRepository) CreateProductImage(prodImage models.ProductImage) error {
	err := pir.database.SetupDatabase().Create(&prodImage).Error
	return err
}
