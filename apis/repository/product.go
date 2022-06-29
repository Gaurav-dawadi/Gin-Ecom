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

func GetProduct(prod_id string) (*models.Product, error) {
	var product models.Product
	err := infrastructure.SetupDatabase().Find(&product, prod_id).Error
	return &product, err
}

func CreateProduct(prod models.Product) (*models.Product, error) {
	err := infrastructure.SetupDatabase().Create(&prod).Error
	return &prod, err
}

func UpdateProduct(prod_id string, prod models.UpdateProduct) error {
	err := infrastructure.SetupDatabase().Model(models.Product{}).
		Where("id = ?", prod_id).
		Updates(models.Product{Name: prod.Name, Description: prod.Description, Quantity: prod.Quantity, Price: prod.Price}).Error

	return err
}

func DeleteProduct(prod_id uint) error {
	return infrastructure.SetupDatabase().
		Where("id = ?", prod_id).
		Delete(&models.Product{}).
		Error
}
