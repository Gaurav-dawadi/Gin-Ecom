package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ps ProductService) GetAllProduct() ([]models.Product, error) {
	return ps.productRepository.GetAllProduct()
}

func (ps ProductService) GetProduct(prod_id string) (*models.Product, error) {
	return ps.productRepository.GetProduct(prod_id)
}

func (ps ProductService) CreateProduct(prod models.Product) (*models.Product, error) {
	prod_re, err := ps.productRepository.CreateProduct(prod)
	return prod_re, err
}

func (ps ProductService) UpdateProduct(prod_id string, prod models.UpdateProduct) error {
	return ps.productRepository.UpdateProduct(prod_id, prod)
}

func (ps ProductService) DeleteProduct(prod_id uint) error {
	return ps.productRepository.DeleteProduct(prod_id)
}
