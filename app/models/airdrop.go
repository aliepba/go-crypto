package models

import (
	"time"

	"github.com/google/uuid"
)

type Airdrop struct {
	ID          uuid.UUID
	ProjectName string
	Description string
	Status      string
	StartDate   time.Time
	EndDate     time.Time
	TotalPrize  int
	WinnerCount int
	CoinID      int
	Coin        Coin
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
