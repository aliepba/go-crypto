package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodMetadata interface {
	FindAllMetadata() ([]models.MetadataCoin, error)
	CreateMetadata(metadata models.MetadataCoin) (models.MetadataCoin, error)
}

func NewMethodMetadata(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) FindAllMetadata() ([]models.MetadataCoin, error) {
	var metadata []models.MetadataCoin
	err := m.db.Find(&metadata).Error

	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func (m *method) CreateMetadata(metadata models.MetadataCoin) (models.MetadataCoin, error) {
	err := m.db.Create(&metadata).Error

	if err != nil {
		return metadata, err
	}

	return metadata, nil
}
