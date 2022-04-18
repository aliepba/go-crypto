package requests

import (
	"time"

	"github.com/google/uuid"
)

type AirdropInput struct {
	ID          uuid.UUID `json:"id"`
	ProjectName string    `json:"project_name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	TotalPrize  int       `json:"total_prize"`
	WinnerCount int       `json:"winner_count"`
	CoinId      int       `json:"coin_id"`
}

type GetAirdropDetail struct {
	ID string `uri:"id"`
}
