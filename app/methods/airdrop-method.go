package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodAirdrop interface {
	FindAllAirdrop() ([]models.Airdrop, error)
	DetailAirdrop(ID string) (models.Airdrop, error)
	CreateAirdrop(airdrop models.Airdrop) (models.Airdrop, error)
	UpdateStatus(airdrop models.Airdrop) (models.Airdrop, error)
}

func NewMethodAirdrop(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) CreateAirdrop(airdrop models.Airdrop) (models.Airdrop, error) {
	err := m.db.Create(&airdrop).Error

	if err != nil {
		return airdrop, err
	}

	return airdrop, nil
}

func (m *method) FindAllAirdrop() ([]models.Airdrop, error) {
	var airdrops []models.Airdrop

	err := m.db.Find(&airdrops).Error

	if err != nil {
		return airdrops, err
	}

	return airdrops, nil
}

func (m *method) DetailAirdrop(ID string) (models.Airdrop, error) {
	var airdrop models.Airdrop
	err := m.db.Preload("Coin").Where("id = ?", ID).Find(&airdrop).Error

	if err != nil {
		return airdrop, err
	}

	return airdrop, err
}

func (m *method) UpdateStatus(airdrop models.Airdrop) (models.Airdrop, error) {
	err := m.db.Save(&airdrop).Error
	if err != nil {
		return airdrop, err
	}

	return airdrop, nil
}
