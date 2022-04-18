package helpers

import (
	"time"

	"github.com/aliepba/go-crypto/app/models"
	"github.com/google/uuid"
)

type AirdropFormatter struct {
	ID          uuid.UUID `json:"id"`
	ProjectName string    `json:"project_name"`
	Status      string    `json:"status"`
	TotalPrize  int       `json:"total_prize"`
	WinnerCount int       `json:"winner_count"`
}

type AirdropDetailFormatter struct {
	ID          uuid.UUID            `json:"id"`
	ProjectName string               `json:"project_name"`
	Description string               `json:"description"`
	Status      string               `json:"status"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	TotalPrize  int                  `json:"total_prize"`
	WinnerCount int                  `json:"winner_count"`
	Coin        CoinAirdropFormatter `json:"coin"`
}

type CoinAirdropFormatter struct {
	Coin string `json:"coin"`
}

func FormatAidrop(airdrop models.Airdrop) AirdropFormatter {
	formatter := AirdropFormatter{}
	formatter.ID = airdrop.ID
	formatter.ProjectName = airdrop.ProjectName
	formatter.Status = airdrop.Status
	formatter.TotalPrize = airdrop.TotalPrize
	formatter.WinnerCount = airdrop.WinnerCount

	return formatter
}

func FormatAirdrops(airdrops []models.Airdrop) []AirdropFormatter {
	airdropsFormatter := []AirdropFormatter{}

	for _, airdrop := range airdrops {
		airdropFormatter := FormatAidrop(airdrop)
		airdropsFormatter = append(airdropsFormatter, airdropFormatter)
	}

	return airdropsFormatter
}

func FormatDetailAirdrop(airdrop models.Airdrop) AirdropDetailFormatter {
	airdropDetailFormatter := AirdropDetailFormatter{}
	airdropDetailFormatter.ID = airdrop.ID
	airdropDetailFormatter.ProjectName = airdrop.ProjectName
	airdropDetailFormatter.Description = airdrop.Description
	airdropDetailFormatter.Status = airdrop.Status
	airdropDetailFormatter.StartDate = airdrop.StartDate
	airdropDetailFormatter.EndDate = airdrop.EndDate
	airdropDetailFormatter.TotalPrize = airdrop.TotalPrize
	airdropDetailFormatter.WinnerCount = airdrop.WinnerCount

	coin := airdrop.Coin
	coinAirdropFormatter := CoinAirdropFormatter{}
	coinAirdropFormatter.Coin = coin.Coin

	airdropDetailFormatter.Coin = coinAirdropFormatter

	return airdropDetailFormatter
}
