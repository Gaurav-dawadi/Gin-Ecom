package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs CategoryService) GetAllCategory() ([]models.Category, error) {
	return cs.categoryRepository.GetAllCategory()
}

func (cs CategoryService) CreateCategory(category models.Category) error {
	err := cs.categoryRepository.CreateCategory(category)
	return err
}
