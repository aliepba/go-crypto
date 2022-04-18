package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
)

type MetadataService interface {
	FindMetadata() ([]models.MetadataCoin, error)
	SaveMetadata(input requests.MetadataInput) (models.MetadataCoin, error)
}

type metadataService struct {
	method methods.MethodMetadata
}

func NewServiceMetadata(method methods.MethodMetadata) *metadataService {
	return &metadataService{method}
}

func (s *metadataService) FindMetadata() ([]models.MetadataCoin, error) {
	metadata, err := s.method.FindAllMetadata()
	if err != nil {
		return metadata, nil
	}

	return metadata, nil
}

func (s *metadataService) SaveMetadata(input requests.MetadataInput) (models.MetadataCoin, error) {
	metadata := models.MetadataCoin{}
	metadata.Website = input.Website
	metadata.TechnicalDoc = input.TechnicalDoc
	metadata.Twitter = input.Twitter
	metadata.Reddit = input.Reddit
	metadata.MessageBoard = input.MessageBoard
	metadata.SourceCode = input.SourceCode
	metadata.CoinID = input.CoinID
	metadata.UserID = int(input.User.ID)

	newMetadata, err := s.method.CreateMetadata(metadata)

	if err != nil {
		return newMetadata, err
	}

	return newMetadata, nil
}
