package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func CreateProductImage(prodId uint, filePath string) error {
	var productImg models.ProductImage

	productImg.ProductID = prodId
	productImg.Image = filePath

	err := repository.CreateProductImage(productImg)
	return err
}
