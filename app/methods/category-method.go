package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodCategory interface {
	FindAllCategory() ([]models.Category, error)
	CoinByCategory(category string) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
}

func NewMethodCategory(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) CreateCategory(category models.Category) (models.Category, error) {
	err := m.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (m *method) FindAllCategory() ([]models.Category, error) {
	var categories []models.Category

	err := m.db.Find(&categories).Error

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (m *method) CoinByCategory(category string) (models.Category, error) {
	var categories models.Category

	err := m.db.Preload("Coin").Where("category = ?", category).Find(&categories).Error

	if err != nil {
		return categories, err
	}

	return categories, nil
}
