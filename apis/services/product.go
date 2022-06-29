package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func GetAllProduct() ([]models.Product, error) {
	return repository.GetAllProduct()
}

func GetProduct(prod_id string) (*models.Product, error) {
	return repository.GetProduct(prod_id)
}

func CreateProduct(prod models.Product) (*models.Product, error) {
	prod_re, err := repository.CreateProduct(prod)
	return prod_re, err
}

func UpdateProduct(prod_id string, prod models.UpdateProduct) error {
	return repository.UpdateProduct(prod_id, prod)
}

func DeleteProduct(prod_id uint) error {
	return repository.DeleteProduct(prod_id)
}
