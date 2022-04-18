package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
)

type AirdropService interface {
	CreateAirdrop(input requests.AirdropInput) (models.Airdrop, error)
	FindAirdrop() ([]models.Airdrop, error)
	FindAirdropByID(input requests.GetAirdropDetail) (models.Airdrop, error)
	UpdateStatus(inputID requests.GetAirdropDetail, input requests.AirdropInput) (models.Airdrop, error)
}

type airdropService struct {
	method methods.MethodAirdrop
}

func NewServiceAirdrop(method methods.MethodAirdrop) *airdropService {
	return &airdropService{method}
}

func (s *airdropService) CreateAirdrop(input requests.AirdropInput) (models.Airdrop, error) {
	airdrop := models.Airdrop{}
	airdrop.ID = input.ID
	airdrop.ProjectName = input.ProjectName
	airdrop.Description = input.Description
	airdrop.StartDate = input.StartDate
	airdrop.EndDate = input.EndDate
	airdrop.TotalPrize = input.TotalPrize
	airdrop.WinnerCount = input.WinnerCount
	airdrop.CoinID = input.CoinId
	airdrop.Status = "UPCOMING"

	newAirdrop, err := s.method.CreateAirdrop(airdrop)

	if err != nil {
		return newAirdrop, err
	}

	return newAirdrop, err
}

func (s *airdropService) FindAirdrop() ([]models.Airdrop, error) {
	airdrops, err := s.method.FindAllAirdrop()

	if err != nil {
		return airdrops, nil
	}

	return airdrops, nil
}

func (s *airdropService) FindAirdropByID(input requests.GetAirdropDetail) (models.Airdrop, error) {
	airdrop, err := s.method.DetailAirdrop(input.ID)

	if err != nil {
		return airdrop, err
	}

	return airdrop, nil

}

func (s *airdropService) UpdateStatus(inputID requests.GetAirdropDetail, input requests.AirdropInput) (models.Airdrop, error) {
	airdrop, err := s.method.DetailAirdrop(inputID.ID)

	if err != nil {
		return airdrop, nil
	}

	airdrop.ID = input.ID
	airdrop.ProjectName = input.ProjectName
	airdrop.Description = input.Description
	airdrop.StartDate = input.StartDate
	airdrop.EndDate = input.EndDate
	airdrop.TotalPrize = input.TotalPrize
	airdrop.WinnerCount = input.WinnerCount
	airdrop.CoinID = input.CoinId
	airdrop.Status = input.Status

	updateAirdrop, err := s.method.UpdateStatus(airdrop)

	if err != nil {
		return updateAirdrop, err
	}

	return updateAirdrop, nil
}
