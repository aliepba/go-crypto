package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodCoin interface {
	FindAll() ([]models.Coin, error)
	CreateCoin(coin models.Coin) (models.Coin, error)
	DetailCoin(symbol string) (models.Coin, error)
}

func NewMethodCoin(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) FindAll() ([]models.Coin, error) {
	var coins []models.Coin
	err := m.db.Preload("MetadataCoin").Find(&coins).Error

	if err != nil {
		return coins, err
	}

	return coins, nil
}

func (m *method) CreateCoin(coin models.Coin) (models.Coin, error) {
	err := m.db.Create(&coin).Error

	if err != nil {
		return coin, err
	}

	return coin, nil
}

func (m *method) DetailCoin(symbol string) (models.Coin, error) {
	var coin models.Coin
	err := m.db.Preload("MetadataCoin").Preload("Category").Where("symbol = ?", symbol).Find(&coin).Error

	if err != nil {
		return coin, err
	}

	return coin, err
}
