package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

type ProductRepository struct {
	database infrastructure.DatabaseSetup
}

func NewProductRepository(database infrastructure.DatabaseSetup) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (pr ProductRepository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := pr.database.SetupDatabase().Find(&products).Error
	return products, err
}

func (pr ProductRepository) GetProduct(prod_id string) (*models.Product, error) {
	var product models.Product
	err := pr.database.SetupDatabase().Find(&product, prod_id).Error
	return &product, err
}

func (pr ProductRepository) CreateProduct(prod models.Product) (*models.Product, error) {
	err := pr.database.SetupDatabase().Create(&prod).Error
	return &prod, err
}

func (pr ProductRepository) UpdateProduct(prod_id string, prod models.UpdateProduct) error {
	err := pr.database.SetupDatabase().Model(models.Product{}).
		Where("id = ?", prod_id).
		Updates(models.Product{Name: prod.Name, Description: prod.Description, Quantity: prod.Quantity, Price: prod.Price}).Error

	return err
}

func (pr ProductRepository) DeleteProduct(prod_id uint) error {
	return pr.database.SetupDatabase().
		Where("id = ?", prod_id).
		Delete(&models.Product{}).
		Error
}
