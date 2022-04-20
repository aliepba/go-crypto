package requests

type FiatInput struct {
	Name   string `json:"name" binding:"required"`
	Sign   string `json:"sign" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`
}
