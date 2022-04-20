package services

import (
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
)

type FiatService interface {
	SaveFiat(input requests.FiatInput) (models.Fiat, error)
}

type fiatService struct {
	method methods.MethodFiat
}

func NewServiceFiat(method methods.MethodFiat) *fiatService {
	return &fiatService{method}
}

func (s *fiatService) SaveFiat(input requests.FiatInput) (models.Fiat, error) {
	fiat := models.Fiat{}
	fiat.Name = input.Name
	fiat.Sign = input.Sign
	fiat.Symbol = input.Symbol

	newFiat, err := s.method.CreateFiat(fiat)

	if err != nil {
		return newFiat, err
	}

	return newFiat, nil
}
