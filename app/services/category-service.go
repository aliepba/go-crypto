package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
)

type CategoryService interface {
	FindCategory() ([]models.Category, error)
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
