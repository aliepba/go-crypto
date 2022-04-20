package helpers

import "github.com/aliepba/go-crypto/app/models"

type FiatFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Symbol string `json:"symbol"`
}

func FormatFiat(fiat models.Fiat) FiatFormatter {
	formatter := FiatFormatter{}
	formatter.ID = int(fiat.ID)
	formatter.Name = fiat.Name
	formatter.Sign = fiat.Sign
	formatter.Symbol = fiat.Symbol

	return formatter
}
