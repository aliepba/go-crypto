package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
)

type CategoryService interface {
	SaveCategory(input requests.CategoryInput) (models.Category, error)
	FindCategory() ([]models.Category, error)
	GetCoinByCategory(input requests.GetCategoryInput) (models.Category, error)
}

type categoryService struct {
	method methods.MethodCategory
}

func NewServiceCategory(method methods.MethodCategory) *categoryService {
	return &categoryService{method}
}

func (s *categoryService) FindCategory() ([]models.Category, error) {
	categories, err := s.method.FindAllCategory()

	if err != nil {
		return categories, nil
	}

	return categories, nil
}

func (s *categoryService) GetCoinByCategory(input requests.GetCategoryInput) (models.Category, error) {
	categories, err := s.method.CoinByCategory(input.Category)

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *categoryService) SaveCategory(input requests.CategoryInput) (models.Category, error) {
	category := models.Category{}
	category.Category = input.Category

	newCategory, err := s.method.CreateCategory(category)

	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}
