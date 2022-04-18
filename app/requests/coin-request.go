package requests

import "github.com/aliepba/go-crypto/app/models"

type CoinInput struct {
	Coin        string `form:"coin" binding:"required"`
	Symbol      string `form:"symbol" binding:"required"`
	Description string `form:"description" binding:"required"`
	User        models.User
}

type GetCoinDetailInput struct {
	Symbol string `uri:"symbol"`
}
