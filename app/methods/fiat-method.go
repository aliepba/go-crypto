package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodFiat interface {
	CreateFiat(fiat models.Fiat) (models.Fiat, error)
}

func NewMethodFiat(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) CreateFiat(fiat models.Fiat) (models.Fiat, error) {
	err := m.db.Create(&fiat).Error

	if err != nil {
		return fiat, err
	}

	return fiat, nil
}
