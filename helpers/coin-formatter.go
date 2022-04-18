package helpers

import "github.com/aliepba/go-crypto/app/models"

type CoinFormatter struct {
	Coin        string `json:"coin"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type CoinDetailFormatter struct {
	Coin         string                  `json:"coin"`
	Symbol       string                  `json:"symbol"`
	Description  string                  `json:"description"`
	Logo         string                  `json:"logo"`
	Category     CategoryDetailFormatter `json:"category"`
	MetadataCoin MetadataDetailFormatter `json:"metadata_coin"`
}

type MetadataDetailFormatter struct {
	Website      string `json:"website"`
	TechnicalDoc string `json:"technical_doc"`
	Twitter      string `json:"twitter"`
	MessageBoard string `json:"message_board"`
	SourceCode   string `json:"source_code"`
}

type CategoryDetailFormatter struct {
	Category string `json:"category"`
}

func FormatCoin(coin models.Coin) CoinFormatter {
	formatter := CoinFormatter{}

	formatter.Coin = coin.Coin
	formatter.Symbol = coin.Symbol
	formatter.Description = coin.Description
	formatter.Logo = coin.Logo

	return formatter

}

func FormatCoins(coins []models.Coin) []CoinFormatter {
	coinsFormatter := []CoinFormatter{}

	for _, coin := range coins {
		coinFormatter := FormatCoin(coin)
		coinsFormatter = append(coinsFormatter, coinFormatter)
	}

	return coinsFormatter
}

func FormatDetailCoin(coin models.Coin) CoinDetailFormatter {
	coinDetailFormatter := CoinDetailFormatter{}
	coinDetailFormatter.Coin = coin.Coin
	coinDetailFormatter.Symbol = coin.Symbol
	coinDetailFormatter.Description = coin.Description
	coinDetailFormatter.Logo = coin.Logo

	metadata := coin.MetadataCoin
	metadataDetailFormatter := MetadataDetailFormatter{}
	metadataDetailFormatter.Website = metadata.Website
	metadataDetailFormatter.TechnicalDoc = metadata.TechnicalDoc
	metadataDetailFormatter.Twitter = metadata.Twitter
	metadataDetailFormatter.MessageBoard = metadata.MessageBoard
	metadataDetailFormatter.SourceCode = metadata.SourceCode

	category := coin.Category
	categoryDetailFormatter := CategoryDetailFormatter{}
	categoryDetailFormatter.Category = category.Category

	coinDetailFormatter.Category = categoryDetailFormatter
	coinDetailFormatter.MetadataCoin = metadataDetailFormatter

	return coinDetailFormatter
}
