package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func ServiceAllProduct() ([]models.Product, error) {
	return repository.GetAllProduct()
}

func CreateProduct(prod models.Product) error {
	err := repository.CreateProduct(prod)
	return err
}
