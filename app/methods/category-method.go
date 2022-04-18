package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodCategory interface {
	FindAllCategory() ([]models.Category, error)
}

func NewMethodCategory(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) FindAllCategory() ([]models.Category, error) {
	var categories []models.Category

	err := m.db.Find(&categories).Error

	if err != nil {
		return categories, err
	}

	return categories, nil
}
