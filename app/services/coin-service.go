package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
)

type CoinService interface {
	FindCoins() ([]models.Coin, error)
	SaveCoin(input requests.CoinInput, fileLocation string) (models.Coin, error)
	GetCoinBySymbol(input requests.GetCoinDetailInput) (models.Coin, error)
}

type coinService struct {
	method methods.MethodCoin
}

func NewServiceCoin(method methods.MethodCoin) *coinService {
	return &coinService{method}
}

func (s *coinService) FindCoins() ([]models.Coin, error) {
	coins, err := s.method.FindAll()
	if err != nil {
		return coins, nil
	}

	return coins, nil
}

func (s *coinService) GetCoinBySymbol(input requests.GetCoinDetailInput) (models.Coin, error) {
	coin, err := s.method.DetailCoin(input.Symbol)

	if err != nil {
		return coin, err
	}

	return coin, nil
}

func (s *coinService) SaveCoin(input requests.CoinInput, fileLocation string) (models.Coin, error) {
	coin := models.Coin{}
	coin.Coin = input.Coin
	coin.Symbol = input.Symbol
	coin.Description = input.Description
	coin.Logo = fileLocation
	coin.UserID = int(input.User.ID)

	newCoin, err := s.method.CreateCoin(coin)

	if err != nil {
		return newCoin, err
	}

	return newCoin, nil
}
