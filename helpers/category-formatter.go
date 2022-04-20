package helpers

import "github.com/aliepba/go-crypto/app/models"

type CategoryFormatter struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type CoinByCategory struct {
	ID       int                     `json:"id"`
	Category string                  `json:"category"`
	Coin     []CoinCategoryFormatter `json:"coin"`
}

type CoinCategoryFormatter struct {
	Coin   string `json:"coin"`
	Symbol string `json:"symbol"`
}

func FormatCategory(category models.Category) CategoryFormatter {
	formatter := CategoryFormatter{}
	formatter.ID = int(category.ID)
	formatter.Category = category.Category

	return formatter
}

func FormatCategories(categories []models.Category) []CategoryFormatter {
	categoriesFormatter := []CategoryFormatter{}

	for _, category := range categories {
		categoryFormatter := FormatCategory(category)
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}

	return categoriesFormatter
}

func FormatByCategory(category models.Category) CoinByCategory {
	categoryCoinFormatter := CoinByCategory{}
	categoryCoinFormatter.ID = int(category.ID)
	categoryCoinFormatter.Category = category.Category

	coins := []CoinCategoryFormatter{}

	for _, coin := range category.Coin {
		coinCategoryFormatter := CoinCategoryFormatter{}
		coinCategoryFormatter.Coin = coin.Coin
		coinCategoryFormatter.Symbol = coin.Symbol
		coins = append(coins, coinCategoryFormatter)
	}

	categoryCoinFormatter.Coin = coins

	return categoryCoinFormatter
}
