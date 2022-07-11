package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

type ProductImageService struct {
	productImageRepository repository.ProductImageRepository
}

func NewProductImageService(productImageRepository repository.ProductImageRepository) *ProductImageService {
	return &ProductImageService{
		productImageRepository: productImageRepository,
	}
}

func (pis ProductImageService) CreateProductImage(prodId uint, filePath string) error {
	var productImg models.ProductImage

	productImg.ProductID = prodId
	productImg.Image = filePath

	err := pis.productImageRepository.CreateProductImage(productImg)
	return err
}
