package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func ServiceAllProduct() ([]models.Product, error) {
	return repository.GetAllProduct()
}

func CreateProduct(prod models.Product) (*models.Product, error) {
	prod_re, err := repository.CreateProduct(prod)
	return prod_re, err
}
